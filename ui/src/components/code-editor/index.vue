<template>
  <div v-show="isShowPlaceholder && placeholder" class="placeholderBox">
    <div class="placeholderLine" v-for="(content, number) in placeholder" :key="number" @click="handlePlaceholderClick">
      <div class="lineNumber">{{ number + 1 }}</div>
      <div class="lineContent">{{ content }}</div>
    </div>
  </div>
  <div class="code-editor-wrapper">
    <section v-show="!isShowPlaceholder || !placeholder" class="code-editor-content" ref="codeEditorRef"></section>
    <div v-if="errorMessage" ref="errorMsgRef" class="error-msg-container">
      <span class="error-icon">
        <Close />
      </span>
      <span class="message">{{ errorMessage }}</span>
    </div>
  </div>
</template>
<script setup lang="ts">
  import { ref, watch, onMounted, nextTick, computed } from 'vue';
  import * as monaco from 'monaco-editor';
  import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker.js?worker';
  import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker.js?worker';
  import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker.js?worker';
  import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker.js?worker';
  import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker.js?worker';
  import { Close } from 'bkui-vue/lib/icon';
  import { IVariableEditParams } from '../../../types/variable';
  import useEditorVariableReplace from '../../utils/hooks/use-editor-variable-replace';
  import { useRoute } from 'vue-router';
  import { getUnReleasedAppVariables, getVariableList } from '../../api/variable';
  import { validateXML, validateYAML, validateJSON } from '../../utils/kv-validate';

  interface errorLineItem {
    lineNumber: number;
    errorInfo: string;
  }

  self.MonacoEnvironment = {
    getWorker(_, label) {
      if (label === 'json') {
        return new jsonWorker();
      }
      if (label === 'css' || label === 'scss' || label === 'less') {
        return new cssWorker();
      }
      if (label === 'html' || label === 'handlebars' || label === 'razor') {
        return new htmlWorker();
      }
      if (label === 'typescript' || label === 'javascript') {
        return new tsWorker();
      }
      return new editorWorker();
    },
  };

  const props = withDefaults(
    defineProps<{
      modelValue: string;
      lfEol?: boolean;
      variables?: IVariableEditParams[];
      editable?: boolean;
      language?: string;
      errorLine?: errorLineItem[];
      placeholder?: string[];
      lineNumbers?: 'on' | 'off' | 'relative' | 'interval';
      minimap?: boolean;
      verticalScrollbarSize?: number;
      horizonScrollbarSize?: number;
      renderLineHighlight?: 'none' | 'gutter' | 'line' | 'all' | undefined;
      renderIndentGuides?: boolean;
      folding?: boolean;
      alwaysConsumeMouseWheel?: boolean;
      contextmenu?: boolean;
      fileEditor?: boolean;
    }>(),
    {
      variables: () => [],
      editable: true,
      lfEol: true,
      language: '',
      lineNumbers: 'on',
      minimap: true,
      verticalScrollbarSize: 10,
      horizonScrollbarSize: 10,
      renderLineHighlight: 'all',
      renderIndentGuides: true,
      folding: true,
      alwaysConsumeMouseWheel: true,
      contextmenu: true,
      fileEditor: true,
    },
  );

  const emit = defineEmits(['update:modelValue', 'change', 'enter', 'paste']);

  const codeEditorRef = ref();
  let editor: monaco.editor.IStandaloneCodeEditor;
  let editorHoverProvider: monaco.IDisposable;
  let editorVariableProvide: monaco.IDisposable;
  const localVal = ref(props.modelValue);
  const route = useRoute();
  const bkBizId = ref(String(route.params.spaceId));
  const appId = ref(Number(route.params.appId));
  const variableNameList = ref<string[]>(['']);
  const privateVariableNameList = ref<string[]>(['']);
  const isShowPlaceholder = ref(true);
  const errorMessage = ref('');
  const errorMsgRef = ref();

  watch(
    () => props.modelValue,
    (val) => {
      if (val !== localVal.value) {
        editor.setValue(val);
        editorHoverProvider = useEditorVariableReplace(editor, props.variables);
      }
    },
  );

  watch(
    () => props.language,
    (val) => {
      // 设置编辑器语言
      monaco.editor.setModelLanguage(editor.getModel() as monaco.editor.ITextModel, val);
      // 清空错误行
      monaco.editor.setModelMarkers(editor.getModel() as monaco.editor.ITextModel, 'error', []);
      // 切换缩进空格数
      editor.getModel()!.updateOptions({ tabSize: tabSize.value });
      // 校验编辑器内容
      if (localVal.value) validate(localVal.value);
    },
  );

  watch(
    () => props.editable,
    (val) => {
      editor.updateOptions({ readOnly: !val });
    },
  );

  watch(
    () => props.variables,
    (val) => {
      if (Array.isArray(val) && val.length > 0) {
        editorHoverProvider = useEditorVariableReplace(editor, val);
      }
    },
  );

  watch(
    () => props.errorLine,
    () => {
      setErrorLine();
    },
  );

  watch(
    () => errorMessage.value,
    (newVal, oldVal) => {
      nextTick(() => {
        if (newVal !== oldVal) {
          if (newVal) {
            if (errorMsgRef.value) {
              const errorMsgHeight = errorMsgRef.value.clientHeight;
              codeEditorRef.value.style.height = `calc(100% - ${errorMsgHeight}px)`;
            }
          } else {
            codeEditorRef.value.style.height = '100%';
          }
        }
      });
    },
  );

  const tabSize = computed(() => {
    if (props.language === 'xml' || props.language === 'yaml') return 2;
    return 4;
  });

  onMounted(() => {
    if (props.fileEditor) {
      handleVariableList();
      autoCompletion();
    }
    if (!editor) {
      registerLanguage();
      editor = monaco.editor.create(codeEditorRef.value as HTMLElement, {
        value: localVal.value,
        theme: 'custom-theme',
        automaticLayout: true,
        language: props.language || 'custom-language',
        readOnly: !props.editable,
        scrollBeyondLastLine: false,
        tabSize: tabSize.value,
        unicodeHighlight: {
          ambiguousCharacters: false,
        },
        lineNumbers: props.lineNumbers,
        minimap: {
          enabled: props.minimap,
        },
        scrollbar: {
          verticalScrollbarSize: props.verticalScrollbarSize,
          horizontalScrollbarSize: props.horizonScrollbarSize,
          alwaysConsumeMouseWheel: props.alwaysConsumeMouseWheel,
        },
        renderLineHighlight: props.renderLineHighlight,
        // @ts-ignore
        renderIndentGuides: props.renderIndentGuides,
        folding: props.folding,
        contextmenu: props.contextmenu,
      });
    }
    if (props.lfEol) {
      const model = editor.getModel() as monaco.editor.ITextModel;
      model.setEOL(monaco.editor.EndOfLineSequence.LF);
    }
    if (Array.isArray(props.variables) && props.variables.length > 0) {
      editorHoverProvider = useEditorVariableReplace(editor, props.variables);
    }

    editor.onDidChangeModelContent(() => {
      localVal.value = editor.getValue();
      validate(localVal.value);
      emit('update:modelValue', localVal.value);
      emit('change', localVal.value);
    });
    // 监听第一次回车事件
    const listener = editor.onKeyDown((event) => {
      if (event.keyCode === monaco.KeyCode.Enter) {
        emit('enter');
        // 取消监听键盘事件
        listener.dispose();
      }
    });
    // 监听第一次粘贴事件
    const pasteListener = editor.onDidPaste(() => {
      emit('paste');
      pasteListener.dispose();
    });
    // 自动换行
    editor.updateOptions({ wordWrap: 'on' });
    editor.onDidBlurEditorWidget(() => {
      // 当编辑器失去焦点时触发的自定义事件处理逻辑
      if (!props.modelValue) {
        isShowPlaceholder.value = true;
      }
    });
  });

  // 添加错误行
  const setErrorLine = () => {
    // 创建错误标记列表
    const markers = props.errorLine!.map(({ lineNumber, errorInfo }) => ({
      startLineNumber: lineNumber,
      endLineNumber: lineNumber,
      startColumn: 1,
      endColumn: 200,
      message: errorInfo,
      severity: monaco.MarkerSeverity.Error,
    }));
    // 设置编辑器模型的标记
    monaco.editor.setModelMarkers(editor.getModel() as monaco.editor.ITextModel, 'error', markers);
  };

  // 获取全局变量和私有变量列表
  const handleVariableList = async () => {
    const variableList = await getVariableList(bkBizId.value, { start: 0, limit: 1000 });
    variableNameList.value = variableList.details.map((item: any) => ` .${item.spec.name} `);
    if (appId.value) {
      const privateVariableList = await getUnReleasedAppVariables(bkBizId.value, appId.value);
      privateVariableNameList.value = privateVariableList.details.map((item: any) => ` .${item.name} `);
      variableNameList.value!.filter((item) => !privateVariableNameList.value!.includes(item));
    }
  };

  // 注册自定义语言
  const registerLanguage = () => {
    // 注册自定义语言
    monaco.languages.register({ id: 'custom-language' });
    monaco.languages.setMonarchTokensProvider('custom-language', {
      tokenizer: {
        root: [{ regex: /\{\{/, action: { token: 'delimiter.curly', next: '@curly' } }],
        curly: [
          { regex: /\}\}/, action: { token: 'delimiter.curly', next: '@pop' } },
          { regex: /[^{}\s]+/, action: 'custom-token' }, // 自定义内部内容的高亮规则
          { regex: /\s+/, action: '' }, // 忽略空白符
        ],
      },
    });
    monaco.editor.defineTheme('custom-theme', {
      base: 'vs-dark',
      inherit: true,
      colors: {},
      rules: [
        { token: 'custom-token', foreground: '38C197' },
        { token: 'delimiter.curly', foreground: '38C197' },
      ],
    });
    // 设置语言配置
    monaco.languages.setLanguageConfiguration('custom-language', {
      brackets: [['{', '}']],
    });
  };

  // 联想输入
  const autoCompletion = () => {
    editorVariableProvide = monaco.languages.registerCompletionItemProvider(props.language || 'custom-language', {
      triggerCharacters: ['{'], // 触发自动补全的字符
      provideCompletionItems(model: any, position: any) {
        const textBeforePosition = model.getValueInRange({
          startLineNumber: position.lineNumber,
          startColumn: 1,
          endLineNumber: position.lineNumber,
          endColumn: position.column,
        });
        // 根据当前的文本内容和光标位置，返回自动补全的候选项列表
        const variableSuggestions = variableNameList.value!.map((item: string) => ({
          label: item, // 候选项的显示文本
          kind: monaco.languages.CompletionItemKind.Variable, // 候选项的类型
          insertText: item, // 插入光标后的文本
          range: new monaco.Range(position.lineNumber, position.column, position.lineNumber, position.column),
        }));
        const suggestions = [...variableSuggestions];
        if (privateVariableNameList.value?.length > 0) {
          const privateVariableSuggestions = privateVariableNameList.value!.map((item: string) => ({
            label: item,
            kind: monaco.languages.CompletionItemKind.Variable,
            insertText: item,
            range: new monaco.Range(position.lineNumber, position.column, position.lineNumber, position.column),
          }));
          suggestions.push(...privateVariableSuggestions);
        }

        if (textBeforePosition === '{{') {
          return {
            suggestions,
          };
        }
        return {
          suggestions: [],
        };
      },
      resolveCompletionItem: (item: any) => item,
    });
  };
  // 打开搜索框
  const openSearch = () => {
    const findAction = editor.getAction('actions.find');
    if (!props.modelValue) handlePlaceholderClick();
    findAction.run();
  };

  const handlePlaceholderClick = () => {
    isShowPlaceholder.value = false;
    nextTick(() => editor.focus());
  };

  // 校验xml、yaml、json数据类型
  const validate = (val: string) => {
    const markers = getValidateDetail(val);

    if (markers.length > 0) {
      const { startLineNumber, startColumn, message } = markers[0];
      errorMessage.value =
        props.language === 'json' ? message : `${message} at line ${startLineNumber}, column ${startColumn}`;
    } else {
      errorMessage.value = '';
    }

    // 编辑器设置错误标记
    monaco.editor.setModelMarkers(editor.getModel() as monaco.editor.ITextModel, 'error', markers);

    return markers.length === 0;
  };

  const getValidateDetail = (val: string) => {
    let markers: any[] = [];

    if (props.language === 'xml') {
      markers = validateXML(val);
    } else if (props.language === 'yaml') {
      markers = validateYAML(val);
    } else if (props.language === 'json') {
      markers = validateJSON(val);
    }

    return markers;
  };

  // @bug vue3的Teleport组件销毁时，子组件的onBeforeUnmount不会被执行，会出现内存泄漏，目前尚未被修复 https://github.com/vuejs/core/issues/6347
  // onBeforeUnmount(() => {
  //   if (editor) {
  //     editor.dispose()
  //   }
  //   if (editorHoverProvider) {
  //     editorHoverProvider.dispose()
  //   }
  // })
  const validateAndMarkErrorLine = (val: string) => {
    const markers = getValidateDetail(val);
    if (markers.length > 0) {
      setLineRevealAndSelected(markers[0].startLineNumber);
    }
    // 编辑器设置错误标记
    monaco.editor.setModelMarkers(editor.getModel() as monaco.editor.ITextModel, 'error', markers);
    return markers.length === 0;
  };

  // 将某一行滚动到编辑器中心并选中
  const setLineRevealAndSelected = (lineNumber: number) => {
    const lineContent = editor.getModel()!.getLineContent(lineNumber);
    const range = new monaco.Range(lineNumber, 1, lineNumber, lineContent.length + 1);
    editor.revealLineInCenter(lineNumber);
    editor.setSelection(range);
  };

  // 返回滚动条顶部
  const scrollToTop = () => {
    editor.revealLineNearTop(0);
  };
  const destroy = () => {
    if (editor) {
      editor.dispose();
    }
    if (editorHoverProvider) {
      editorHoverProvider.dispose();
    }
    if (editorVariableProvide) {
      editorVariableProvide.dispose();
    }
  };

  defineExpose({
    destroy,
    openSearch,
    validate,
    validateAndMarkErrorLine,
    scrollToTop,
  });
</script>
<style lang="scss" scoped>
  .code-editor-wrapper {
    height: 100%;
    .code-editor-content {
      height: 100%;
    }
    :deep(.monaco-editor) {
      width: 100%;
      .template-variable-item {
        color: #1768ef;
        border: 1px solid #1768ef;
        cursor: pointer;
      }
    }
  }
  .placeholderBox {
    height: 100%;
    background-color: #1e1e1e;
    box-sizing: content-box;
    .placeholderLine {
      display: flex;
      height: 19px;
      line-height: 19px;
      .lineNumber {
        font-family: Consolas, 'Courier New', monospace;
        width: 64px;
        text-align: center;
        color: #979ba5;
        font-size: 14px;
      }
      .lineContent {
        color: #63656e;
      }
    }
  }
  .error-msg-container {
    display: flex;
    align-items: flex-start;
    padding: 8px 16px;
    background: #212121;
    border-left: 4px solid #b34747;
    max-height: 60px;
    overflow: auto;
    .error-icon {
      display: flex;
      align-items: center;
      color: #b34747;
      height: 20px;
      font-size: 12px;
    }
    .message {
      margin-left: 8px;
      color: #dcdee5;
      line-height: 20px;
      font-size: 12px;
      word-break: break-all;
    }
  }
</style>
