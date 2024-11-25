// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/TencentBlueKing/bk-bscp/pkg/dal/table"
)

func newReleasedHook(db *gorm.DB, opts ...gen.DOOption) releasedHook {
	_releasedHook := releasedHook{}

	_releasedHook.releasedHookDo.UseDB(db, opts...)
	_releasedHook.releasedHookDo.UseModel(&table.ReleasedHook{})

	tableName := _releasedHook.releasedHookDo.TableName()
	_releasedHook.ALL = field.NewAsterisk(tableName)
	_releasedHook.ID = field.NewUint32(tableName, "id")
	_releasedHook.AppID = field.NewUint32(tableName, "app_id")
	_releasedHook.ReleaseID = field.NewUint32(tableName, "release_id")
	_releasedHook.HookID = field.NewUint32(tableName, "hook_id")
	_releasedHook.HookRevisionID = field.NewUint32(tableName, "hook_revision_id")
	_releasedHook.HookName = field.NewString(tableName, "hook_name")
	_releasedHook.HookRevisionName = field.NewString(tableName, "hook_revision_name")
	_releasedHook.Content = field.NewString(tableName, "content")
	_releasedHook.ScriptType = field.NewString(tableName, "script_type")
	_releasedHook.HookType = field.NewString(tableName, "hook_type")
	_releasedHook.BizID = field.NewUint32(tableName, "biz_id")
	_releasedHook.Reviser = field.NewString(tableName, "reviser")
	_releasedHook.UpdatedAt = field.NewTime(tableName, "updated_at")

	_releasedHook.fillFieldMap()

	return _releasedHook
}

type releasedHook struct {
	releasedHookDo releasedHookDo

	ALL              field.Asterisk
	ID               field.Uint32
	AppID            field.Uint32
	ReleaseID        field.Uint32
	HookID           field.Uint32
	HookRevisionID   field.Uint32
	HookName         field.String
	HookRevisionName field.String
	Content          field.String
	ScriptType       field.String
	HookType         field.String
	BizID            field.Uint32
	Reviser          field.String
	UpdatedAt        field.Time

	fieldMap map[string]field.Expr
}

func (r releasedHook) Table(newTableName string) *releasedHook {
	r.releasedHookDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r releasedHook) As(alias string) *releasedHook {
	r.releasedHookDo.DO = *(r.releasedHookDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *releasedHook) updateTableName(table string) *releasedHook {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint32(table, "id")
	r.AppID = field.NewUint32(table, "app_id")
	r.ReleaseID = field.NewUint32(table, "release_id")
	r.HookID = field.NewUint32(table, "hook_id")
	r.HookRevisionID = field.NewUint32(table, "hook_revision_id")
	r.HookName = field.NewString(table, "hook_name")
	r.HookRevisionName = field.NewString(table, "hook_revision_name")
	r.Content = field.NewString(table, "content")
	r.ScriptType = field.NewString(table, "script_type")
	r.HookType = field.NewString(table, "hook_type")
	r.BizID = field.NewUint32(table, "biz_id")
	r.Reviser = field.NewString(table, "reviser")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *releasedHook) WithContext(ctx context.Context) IReleasedHookDo {
	return r.releasedHookDo.WithContext(ctx)
}

func (r releasedHook) TableName() string { return r.releasedHookDo.TableName() }

func (r releasedHook) Alias() string { return r.releasedHookDo.Alias() }

func (r releasedHook) Columns(cols ...field.Expr) gen.Columns {
	return r.releasedHookDo.Columns(cols...)
}

func (r *releasedHook) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *releasedHook) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 13)
	r.fieldMap["id"] = r.ID
	r.fieldMap["app_id"] = r.AppID
	r.fieldMap["release_id"] = r.ReleaseID
	r.fieldMap["hook_id"] = r.HookID
	r.fieldMap["hook_revision_id"] = r.HookRevisionID
	r.fieldMap["hook_name"] = r.HookName
	r.fieldMap["hook_revision_name"] = r.HookRevisionName
	r.fieldMap["content"] = r.Content
	r.fieldMap["script_type"] = r.ScriptType
	r.fieldMap["hook_type"] = r.HookType
	r.fieldMap["biz_id"] = r.BizID
	r.fieldMap["reviser"] = r.Reviser
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r releasedHook) clone(db *gorm.DB) releasedHook {
	r.releasedHookDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r releasedHook) replaceDB(db *gorm.DB) releasedHook {
	r.releasedHookDo.ReplaceDB(db)
	return r
}

type releasedHookDo struct{ gen.DO }

type IReleasedHookDo interface {
	gen.SubQuery
	Debug() IReleasedHookDo
	WithContext(ctx context.Context) IReleasedHookDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReleasedHookDo
	WriteDB() IReleasedHookDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReleasedHookDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReleasedHookDo
	Not(conds ...gen.Condition) IReleasedHookDo
	Or(conds ...gen.Condition) IReleasedHookDo
	Select(conds ...field.Expr) IReleasedHookDo
	Where(conds ...gen.Condition) IReleasedHookDo
	Order(conds ...field.Expr) IReleasedHookDo
	Distinct(cols ...field.Expr) IReleasedHookDo
	Omit(cols ...field.Expr) IReleasedHookDo
	Join(table schema.Tabler, on ...field.Expr) IReleasedHookDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReleasedHookDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReleasedHookDo
	Group(cols ...field.Expr) IReleasedHookDo
	Having(conds ...gen.Condition) IReleasedHookDo
	Limit(limit int) IReleasedHookDo
	Offset(offset int) IReleasedHookDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReleasedHookDo
	Unscoped() IReleasedHookDo
	Create(values ...*table.ReleasedHook) error
	CreateInBatches(values []*table.ReleasedHook, batchSize int) error
	Save(values ...*table.ReleasedHook) error
	First() (*table.ReleasedHook, error)
	Take() (*table.ReleasedHook, error)
	Last() (*table.ReleasedHook, error)
	Find() ([]*table.ReleasedHook, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ReleasedHook, err error)
	FindInBatches(result *[]*table.ReleasedHook, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.ReleasedHook) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReleasedHookDo
	Assign(attrs ...field.AssignExpr) IReleasedHookDo
	Joins(fields ...field.RelationField) IReleasedHookDo
	Preload(fields ...field.RelationField) IReleasedHookDo
	FirstOrInit() (*table.ReleasedHook, error)
	FirstOrCreate() (*table.ReleasedHook, error)
	FindByPage(offset int, limit int) (result []*table.ReleasedHook, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReleasedHookDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r releasedHookDo) Debug() IReleasedHookDo {
	return r.withDO(r.DO.Debug())
}

func (r releasedHookDo) WithContext(ctx context.Context) IReleasedHookDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r releasedHookDo) ReadDB() IReleasedHookDo {
	return r.Clauses(dbresolver.Read)
}

func (r releasedHookDo) WriteDB() IReleasedHookDo {
	return r.Clauses(dbresolver.Write)
}

func (r releasedHookDo) Session(config *gorm.Session) IReleasedHookDo {
	return r.withDO(r.DO.Session(config))
}

func (r releasedHookDo) Clauses(conds ...clause.Expression) IReleasedHookDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r releasedHookDo) Returning(value interface{}, columns ...string) IReleasedHookDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r releasedHookDo) Not(conds ...gen.Condition) IReleasedHookDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r releasedHookDo) Or(conds ...gen.Condition) IReleasedHookDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r releasedHookDo) Select(conds ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r releasedHookDo) Where(conds ...gen.Condition) IReleasedHookDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r releasedHookDo) Order(conds ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r releasedHookDo) Distinct(cols ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r releasedHookDo) Omit(cols ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r releasedHookDo) Join(table schema.Tabler, on ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r releasedHookDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r releasedHookDo) RightJoin(table schema.Tabler, on ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r releasedHookDo) Group(cols ...field.Expr) IReleasedHookDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r releasedHookDo) Having(conds ...gen.Condition) IReleasedHookDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r releasedHookDo) Limit(limit int) IReleasedHookDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r releasedHookDo) Offset(offset int) IReleasedHookDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r releasedHookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReleasedHookDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r releasedHookDo) Unscoped() IReleasedHookDo {
	return r.withDO(r.DO.Unscoped())
}

func (r releasedHookDo) Create(values ...*table.ReleasedHook) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r releasedHookDo) CreateInBatches(values []*table.ReleasedHook, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r releasedHookDo) Save(values ...*table.ReleasedHook) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r releasedHookDo) First() (*table.ReleasedHook, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedHook), nil
	}
}

func (r releasedHookDo) Take() (*table.ReleasedHook, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedHook), nil
	}
}

func (r releasedHookDo) Last() (*table.ReleasedHook, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedHook), nil
	}
}

func (r releasedHookDo) Find() ([]*table.ReleasedHook, error) {
	result, err := r.DO.Find()
	return result.([]*table.ReleasedHook), err
}

func (r releasedHookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ReleasedHook, err error) {
	buf := make([]*table.ReleasedHook, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r releasedHookDo) FindInBatches(result *[]*table.ReleasedHook, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r releasedHookDo) Attrs(attrs ...field.AssignExpr) IReleasedHookDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r releasedHookDo) Assign(attrs ...field.AssignExpr) IReleasedHookDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r releasedHookDo) Joins(fields ...field.RelationField) IReleasedHookDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r releasedHookDo) Preload(fields ...field.RelationField) IReleasedHookDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r releasedHookDo) FirstOrInit() (*table.ReleasedHook, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedHook), nil
	}
}

func (r releasedHookDo) FirstOrCreate() (*table.ReleasedHook, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedHook), nil
	}
}

func (r releasedHookDo) FindByPage(offset int, limit int) (result []*table.ReleasedHook, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r releasedHookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r releasedHookDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r releasedHookDo) Delete(models ...*table.ReleasedHook) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *releasedHookDo) withDO(do gen.Dao) *releasedHookDo {
	r.DO = *do.(*gen.DO)
	return r
}
