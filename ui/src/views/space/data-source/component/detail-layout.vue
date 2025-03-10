<template>
  <section :class="['data-source-detail-layout', { 'show-notice': showNotice }]">
    <header>
      <div class="nav-title">
        <ArrowsLeft class="arrow-icon" @click="emits('close')" />
        <span class="title">{{ props.name }}</span>
        <template v-if="props.suffix">
          <span class="line"></span>
          <span class="suffix"> {{ props.suffix }}</span>
        </template>
      </div>
    </header>
    <div :class="['layout-content', { 'without-footer': !props.showFooter }]">
      <slot name="content"></slot>
    </div>
    <footer v-if="props.showFooter" class="layout-footer">
      <slot name="footer"></slot>
    </footer>
  </section>
</template>
<script setup lang="ts">
  import { ArrowsLeft } from 'bkui-vue/lib/icon';
  import { storeToRefs } from 'pinia';
  import useGlobalStore from '../../../../store/global';

  const { showNotice } = storeToRefs(useGlobalStore());

  const props = withDefaults(
    defineProps<{
      name: string;
      suffix?: string;
      showFooter?: boolean;
    }>(),
    {
      showFooter: true,
    },
  );

  const emits = defineEmits(['close']);
</script>
<style lang="scss" scoped>
  .data-source-detail-layout {
    position: fixed;
    top: 52px;
    left: 0;
    width: 100%;
    height: calc(100vh - 52px);
    background: #ffffff;
    z-index: 2000;
    &.show-notice {
      top: 92px;
      height: calc(100vh - 92px);
    }
    .nav-title {
      display: flex;
      align-items: center;
      position: relative;
      padding: 0 24px;
      background: #ffffff;
      box-shadow: 0 3px 4px 0 #0000000a;
      z-index: 1;
    }
    .arrow-icon {
      font-size: 24px;
      color: #3a84ff;
      cursor: pointer;
    }
    .title {
      padding: 14px 0;
      font-size: 16px;
      line-height: 24px;
      color: #313238;
    }
    .line {
      width: 1px;
      height: 16px;
      background-color: #dcdee5;
      margin: 0 16px;
    }
    .suffix {
      font-size: 16px;
      color: #63656e;
    }
    .layout-content {
      height: calc(100% - 100px);
      overflow: auto;
      &.without-footer {
        height: calc(100% - 52px);
      }
    }
    .layout-footer {
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 8px 0;
      background: #fafbfd;
      box-shadow: 0 -1px 0 0 #dcdee5;
    }
  }
</style>
