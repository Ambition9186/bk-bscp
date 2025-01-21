// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                           = new(Query)
	App                         *app
	AppTemplateBinding          *appTemplateBinding
	AppTemplateVariable         *appTemplateVariable
	ArchivedApp                 *archivedApp
	Audit                       *audit
	Client                      *client
	ClientEvent                 *clientEvent
	ClientQuery                 *clientQuery
	Commit                      *commit
	Config                      *config
	ConfigItem                  *configItem
	Content                     *content
	Credential                  *credential
	CredentialScope             *credentialScope
	DataSourceContent           *dataSourceContent
	DataSourceInfo              *dataSourceInfo
	DataSourceMapping           *dataSourceMapping
	Event                       *event
	Group                       *group
	GroupAppBind                *groupAppBind
	Hook                        *hook
	HookRevision                *hookRevision
	IDGenerator                 *iDGenerator
	Kv                          *kv
	Release                     *release
	ReleasedAppTemplate         *releasedAppTemplate
	ReleasedAppTemplateVariable *releasedAppTemplateVariable
	ReleasedConfigItem          *releasedConfigItem
	ReleasedGroup               *releasedGroup
	ReleasedHook                *releasedHook
	ReleasedKv                  *releasedKv
	ReleasedTableContent        *releasedTableContent
	ResourceLock                *resourceLock
	Strategy                    *strategy
	Template                    *template
	TemplateRevision            *templateRevision
	TemplateSet                 *templateSet
	TemplateSpace               *templateSpace
	TemplateVariable            *templateVariable
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	App = &Q.App
	AppTemplateBinding = &Q.AppTemplateBinding
	AppTemplateVariable = &Q.AppTemplateVariable
	ArchivedApp = &Q.ArchivedApp
	Audit = &Q.Audit
	Client = &Q.Client
	ClientEvent = &Q.ClientEvent
	ClientQuery = &Q.ClientQuery
	Commit = &Q.Commit
	Config = &Q.Config
	ConfigItem = &Q.ConfigItem
	Content = &Q.Content
	Credential = &Q.Credential
	CredentialScope = &Q.CredentialScope
	DataSourceContent = &Q.DataSourceContent
	DataSourceInfo = &Q.DataSourceInfo
	DataSourceMapping = &Q.DataSourceMapping
	Event = &Q.Event
	Group = &Q.Group
	GroupAppBind = &Q.GroupAppBind
	Hook = &Q.Hook
	HookRevision = &Q.HookRevision
	IDGenerator = &Q.IDGenerator
	Kv = &Q.Kv
	Release = &Q.Release
	ReleasedAppTemplate = &Q.ReleasedAppTemplate
	ReleasedAppTemplateVariable = &Q.ReleasedAppTemplateVariable
	ReleasedConfigItem = &Q.ReleasedConfigItem
	ReleasedGroup = &Q.ReleasedGroup
	ReleasedHook = &Q.ReleasedHook
	ReleasedKv = &Q.ReleasedKv
	ReleasedTableContent = &Q.ReleasedTableContent
	ResourceLock = &Q.ResourceLock
	Strategy = &Q.Strategy
	Template = &Q.Template
	TemplateRevision = &Q.TemplateRevision
	TemplateSet = &Q.TemplateSet
	TemplateSpace = &Q.TemplateSpace
	TemplateVariable = &Q.TemplateVariable
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                          db,
		App:                         newApp(db, opts...),
		AppTemplateBinding:          newAppTemplateBinding(db, opts...),
		AppTemplateVariable:         newAppTemplateVariable(db, opts...),
		ArchivedApp:                 newArchivedApp(db, opts...),
		Audit:                       newAudit(db, opts...),
		Client:                      newClient(db, opts...),
		ClientEvent:                 newClientEvent(db, opts...),
		ClientQuery:                 newClientQuery(db, opts...),
		Commit:                      newCommit(db, opts...),
		Config:                      newConfig(db, opts...),
		ConfigItem:                  newConfigItem(db, opts...),
		Content:                     newContent(db, opts...),
		Credential:                  newCredential(db, opts...),
		CredentialScope:             newCredentialScope(db, opts...),
		DataSourceContent:           newDataSourceContent(db, opts...),
		DataSourceInfo:              newDataSourceInfo(db, opts...),
		DataSourceMapping:           newDataSourceMapping(db, opts...),
		Event:                       newEvent(db, opts...),
		Group:                       newGroup(db, opts...),
		GroupAppBind:                newGroupAppBind(db, opts...),
		Hook:                        newHook(db, opts...),
		HookRevision:                newHookRevision(db, opts...),
		IDGenerator:                 newIDGenerator(db, opts...),
		Kv:                          newKv(db, opts...),
		Release:                     newRelease(db, opts...),
		ReleasedAppTemplate:         newReleasedAppTemplate(db, opts...),
		ReleasedAppTemplateVariable: newReleasedAppTemplateVariable(db, opts...),
		ReleasedConfigItem:          newReleasedConfigItem(db, opts...),
		ReleasedGroup:               newReleasedGroup(db, opts...),
		ReleasedHook:                newReleasedHook(db, opts...),
		ReleasedKv:                  newReleasedKv(db, opts...),
		ReleasedTableContent:        newReleasedTableContent(db, opts...),
		ResourceLock:                newResourceLock(db, opts...),
		Strategy:                    newStrategy(db, opts...),
		Template:                    newTemplate(db, opts...),
		TemplateRevision:            newTemplateRevision(db, opts...),
		TemplateSet:                 newTemplateSet(db, opts...),
		TemplateSpace:               newTemplateSpace(db, opts...),
		TemplateVariable:            newTemplateVariable(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	App                         app
	AppTemplateBinding          appTemplateBinding
	AppTemplateVariable         appTemplateVariable
	ArchivedApp                 archivedApp
	Audit                       audit
	Client                      client
	ClientEvent                 clientEvent
	ClientQuery                 clientQuery
	Commit                      commit
	Config                      config
	ConfigItem                  configItem
	Content                     content
	Credential                  credential
	CredentialScope             credentialScope
	DataSourceContent           dataSourceContent
	DataSourceInfo              dataSourceInfo
	DataSourceMapping           dataSourceMapping
	Event                       event
	Group                       group
	GroupAppBind                groupAppBind
	Hook                        hook
	HookRevision                hookRevision
	IDGenerator                 iDGenerator
	Kv                          kv
	Release                     release
	ReleasedAppTemplate         releasedAppTemplate
	ReleasedAppTemplateVariable releasedAppTemplateVariable
	ReleasedConfigItem          releasedConfigItem
	ReleasedGroup               releasedGroup
	ReleasedHook                releasedHook
	ReleasedKv                  releasedKv
	ReleasedTableContent        releasedTableContent
	ResourceLock                resourceLock
	Strategy                    strategy
	Template                    template
	TemplateRevision            templateRevision
	TemplateSet                 templateSet
	TemplateSpace               templateSpace
	TemplateVariable            templateVariable
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                          db,
		App:                         q.App.clone(db),
		AppTemplateBinding:          q.AppTemplateBinding.clone(db),
		AppTemplateVariable:         q.AppTemplateVariable.clone(db),
		ArchivedApp:                 q.ArchivedApp.clone(db),
		Audit:                       q.Audit.clone(db),
		Client:                      q.Client.clone(db),
		ClientEvent:                 q.ClientEvent.clone(db),
		ClientQuery:                 q.ClientQuery.clone(db),
		Commit:                      q.Commit.clone(db),
		Config:                      q.Config.clone(db),
		ConfigItem:                  q.ConfigItem.clone(db),
		Content:                     q.Content.clone(db),
		Credential:                  q.Credential.clone(db),
		CredentialScope:             q.CredentialScope.clone(db),
		DataSourceContent:           q.DataSourceContent.clone(db),
		DataSourceInfo:              q.DataSourceInfo.clone(db),
		DataSourceMapping:           q.DataSourceMapping.clone(db),
		Event:                       q.Event.clone(db),
		Group:                       q.Group.clone(db),
		GroupAppBind:                q.GroupAppBind.clone(db),
		Hook:                        q.Hook.clone(db),
		HookRevision:                q.HookRevision.clone(db),
		IDGenerator:                 q.IDGenerator.clone(db),
		Kv:                          q.Kv.clone(db),
		Release:                     q.Release.clone(db),
		ReleasedAppTemplate:         q.ReleasedAppTemplate.clone(db),
		ReleasedAppTemplateVariable: q.ReleasedAppTemplateVariable.clone(db),
		ReleasedConfigItem:          q.ReleasedConfigItem.clone(db),
		ReleasedGroup:               q.ReleasedGroup.clone(db),
		ReleasedHook:                q.ReleasedHook.clone(db),
		ReleasedKv:                  q.ReleasedKv.clone(db),
		ReleasedTableContent:        q.ReleasedTableContent.clone(db),
		ResourceLock:                q.ResourceLock.clone(db),
		Strategy:                    q.Strategy.clone(db),
		Template:                    q.Template.clone(db),
		TemplateRevision:            q.TemplateRevision.clone(db),
		TemplateSet:                 q.TemplateSet.clone(db),
		TemplateSpace:               q.TemplateSpace.clone(db),
		TemplateVariable:            q.TemplateVariable.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                          db,
		App:                         q.App.replaceDB(db),
		AppTemplateBinding:          q.AppTemplateBinding.replaceDB(db),
		AppTemplateVariable:         q.AppTemplateVariable.replaceDB(db),
		ArchivedApp:                 q.ArchivedApp.replaceDB(db),
		Audit:                       q.Audit.replaceDB(db),
		Client:                      q.Client.replaceDB(db),
		ClientEvent:                 q.ClientEvent.replaceDB(db),
		ClientQuery:                 q.ClientQuery.replaceDB(db),
		Commit:                      q.Commit.replaceDB(db),
		Config:                      q.Config.replaceDB(db),
		ConfigItem:                  q.ConfigItem.replaceDB(db),
		Content:                     q.Content.replaceDB(db),
		Credential:                  q.Credential.replaceDB(db),
		CredentialScope:             q.CredentialScope.replaceDB(db),
		DataSourceContent:           q.DataSourceContent.replaceDB(db),
		DataSourceInfo:              q.DataSourceInfo.replaceDB(db),
		DataSourceMapping:           q.DataSourceMapping.replaceDB(db),
		Event:                       q.Event.replaceDB(db),
		Group:                       q.Group.replaceDB(db),
		GroupAppBind:                q.GroupAppBind.replaceDB(db),
		Hook:                        q.Hook.replaceDB(db),
		HookRevision:                q.HookRevision.replaceDB(db),
		IDGenerator:                 q.IDGenerator.replaceDB(db),
		Kv:                          q.Kv.replaceDB(db),
		Release:                     q.Release.replaceDB(db),
		ReleasedAppTemplate:         q.ReleasedAppTemplate.replaceDB(db),
		ReleasedAppTemplateVariable: q.ReleasedAppTemplateVariable.replaceDB(db),
		ReleasedConfigItem:          q.ReleasedConfigItem.replaceDB(db),
		ReleasedGroup:               q.ReleasedGroup.replaceDB(db),
		ReleasedHook:                q.ReleasedHook.replaceDB(db),
		ReleasedKv:                  q.ReleasedKv.replaceDB(db),
		ReleasedTableContent:        q.ReleasedTableContent.replaceDB(db),
		ResourceLock:                q.ResourceLock.replaceDB(db),
		Strategy:                    q.Strategy.replaceDB(db),
		Template:                    q.Template.replaceDB(db),
		TemplateRevision:            q.TemplateRevision.replaceDB(db),
		TemplateSet:                 q.TemplateSet.replaceDB(db),
		TemplateSpace:               q.TemplateSpace.replaceDB(db),
		TemplateVariable:            q.TemplateVariable.replaceDB(db),
	}
}

type queryCtx struct {
	App                         IAppDo
	AppTemplateBinding          IAppTemplateBindingDo
	AppTemplateVariable         IAppTemplateVariableDo
	ArchivedApp                 IArchivedAppDo
	Audit                       IAuditDo
	Client                      IClientDo
	ClientEvent                 IClientEventDo
	ClientQuery                 IClientQueryDo
	Commit                      ICommitDo
	Config                      IConfigDo
	ConfigItem                  IConfigItemDo
	Content                     IContentDo
	Credential                  ICredentialDo
	CredentialScope             ICredentialScopeDo
	DataSourceContent           IDataSourceContentDo
	DataSourceInfo              IDataSourceInfoDo
	DataSourceMapping           IDataSourceMappingDo
	Event                       IEventDo
	Group                       IGroupDo
	GroupAppBind                IGroupAppBindDo
	Hook                        IHookDo
	HookRevision                IHookRevisionDo
	IDGenerator                 IIDGeneratorDo
	Kv                          IKvDo
	Release                     IReleaseDo
	ReleasedAppTemplate         IReleasedAppTemplateDo
	ReleasedAppTemplateVariable IReleasedAppTemplateVariableDo
	ReleasedConfigItem          IReleasedConfigItemDo
	ReleasedGroup               IReleasedGroupDo
	ReleasedHook                IReleasedHookDo
	ReleasedKv                  IReleasedKvDo
	ReleasedTableContent        IReleasedTableContentDo
	ResourceLock                IResourceLockDo
	Strategy                    IStrategyDo
	Template                    ITemplateDo
	TemplateRevision            ITemplateRevisionDo
	TemplateSet                 ITemplateSetDo
	TemplateSpace               ITemplateSpaceDo
	TemplateVariable            ITemplateVariableDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		App:                         q.App.WithContext(ctx),
		AppTemplateBinding:          q.AppTemplateBinding.WithContext(ctx),
		AppTemplateVariable:         q.AppTemplateVariable.WithContext(ctx),
		ArchivedApp:                 q.ArchivedApp.WithContext(ctx),
		Audit:                       q.Audit.WithContext(ctx),
		Client:                      q.Client.WithContext(ctx),
		ClientEvent:                 q.ClientEvent.WithContext(ctx),
		ClientQuery:                 q.ClientQuery.WithContext(ctx),
		Commit:                      q.Commit.WithContext(ctx),
		Config:                      q.Config.WithContext(ctx),
		ConfigItem:                  q.ConfigItem.WithContext(ctx),
		Content:                     q.Content.WithContext(ctx),
		Credential:                  q.Credential.WithContext(ctx),
		CredentialScope:             q.CredentialScope.WithContext(ctx),
		DataSourceContent:           q.DataSourceContent.WithContext(ctx),
		DataSourceInfo:              q.DataSourceInfo.WithContext(ctx),
		DataSourceMapping:           q.DataSourceMapping.WithContext(ctx),
		Event:                       q.Event.WithContext(ctx),
		Group:                       q.Group.WithContext(ctx),
		GroupAppBind:                q.GroupAppBind.WithContext(ctx),
		Hook:                        q.Hook.WithContext(ctx),
		HookRevision:                q.HookRevision.WithContext(ctx),
		IDGenerator:                 q.IDGenerator.WithContext(ctx),
		Kv:                          q.Kv.WithContext(ctx),
		Release:                     q.Release.WithContext(ctx),
		ReleasedAppTemplate:         q.ReleasedAppTemplate.WithContext(ctx),
		ReleasedAppTemplateVariable: q.ReleasedAppTemplateVariable.WithContext(ctx),
		ReleasedConfigItem:          q.ReleasedConfigItem.WithContext(ctx),
		ReleasedGroup:               q.ReleasedGroup.WithContext(ctx),
		ReleasedHook:                q.ReleasedHook.WithContext(ctx),
		ReleasedKv:                  q.ReleasedKv.WithContext(ctx),
		ReleasedTableContent:        q.ReleasedTableContent.WithContext(ctx),
		ResourceLock:                q.ResourceLock.WithContext(ctx),
		Strategy:                    q.Strategy.WithContext(ctx),
		Template:                    q.Template.WithContext(ctx),
		TemplateRevision:            q.TemplateRevision.WithContext(ctx),
		TemplateSet:                 q.TemplateSet.WithContext(ctx),
		TemplateSpace:               q.TemplateSpace.WithContext(ctx),
		TemplateVariable:            q.TemplateVariable.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
