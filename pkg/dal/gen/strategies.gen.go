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

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
)

func newStrategy(db *gorm.DB, opts ...gen.DOOption) strategy {
	_strategy := strategy{}

	_strategy.strategyDo.UseDB(db, opts...)
	_strategy.strategyDo.UseModel(&table.Strategy{})

	tableName := _strategy.strategyDo.TableName()
	_strategy.ALL = field.NewAsterisk(tableName)
	_strategy.ID = field.NewUint32(tableName, "id")
	_strategy.Name = field.NewString(tableName, "name")
	_strategy.ReleaseID = field.NewUint32(tableName, "release_id")
	_strategy.AsDefault = field.NewBool(tableName, "as_default")
	_strategy.Scope = field.NewField(tableName, "scope")
	_strategy.Namespace = field.NewString(tableName, "namespace")
	_strategy.Memo = field.NewString(tableName, "memo")
	_strategy.PublishType = field.NewString(tableName, "publish_type")
	_strategy.PublishTime = field.NewString(tableName, "publish_time")
	_strategy.PublishStatus = field.NewString(tableName, "publish_status")
	_strategy.RejectReason = field.NewString(tableName, "reject_reason")
	_strategy.Approver = field.NewString(tableName, "approver")
	_strategy.ApproverProgress = field.NewString(tableName, "approver_progress")
	_strategy.FinalApprovalTime = field.NewTime(tableName, "final_approval_time")
	_strategy.ItsmTicketType = field.NewString(tableName, "itsm_ticket_type")
	_strategy.ItsmTicketUrl = field.NewString(tableName, "itsm_ticket_url")
	_strategy.ItsmTicketSn = field.NewString(tableName, "itsm_ticket_sn")
	_strategy.ItsmTicketStatus = field.NewString(tableName, "itsm_ticket_status")
	_strategy.ItsmTicketStateID = field.NewInt(tableName, "itsm_ticket_state_id")
	_strategy.PubState = field.NewString(tableName, "pub_state")
	_strategy.BizID = field.NewUint32(tableName, "biz_id")
	_strategy.AppID = field.NewUint32(tableName, "app_id")
	_strategy.StrategySetID = field.NewUint32(tableName, "strategy_set_id")
	_strategy.Creator = field.NewString(tableName, "creator")
	_strategy.Reviser = field.NewString(tableName, "reviser")
	_strategy.CreatedAt = field.NewTime(tableName, "created_at")
	_strategy.UpdatedAt = field.NewTime(tableName, "updated_at")

	_strategy.fillFieldMap()

	return _strategy
}

type strategy struct {
	strategyDo strategyDo

	ALL               field.Asterisk
	ID                field.Uint32
	Name              field.String
	ReleaseID         field.Uint32
	AsDefault         field.Bool
	Scope             field.Field
	Namespace         field.String
	Memo              field.String
	PublishType       field.String
	PublishTime       field.String
	PublishStatus     field.String
	RejectReason      field.String
	Approver          field.String
	ApproverProgress  field.String
	FinalApprovalTime field.Time
	ItsmTicketType    field.String
	ItsmTicketUrl     field.String
	ItsmTicketSn      field.String
	ItsmTicketStatus  field.String
	ItsmTicketStateID field.Int
	PubState          field.String
	BizID             field.Uint32
	AppID             field.Uint32
	StrategySetID     field.Uint32
	Creator           field.String
	Reviser           field.String
	CreatedAt         field.Time
	UpdatedAt         field.Time

	fieldMap map[string]field.Expr
}

func (s strategy) Table(newTableName string) *strategy {
	s.strategyDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s strategy) As(alias string) *strategy {
	s.strategyDo.DO = *(s.strategyDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *strategy) updateTableName(table string) *strategy {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint32(table, "id")
	s.Name = field.NewString(table, "name")
	s.ReleaseID = field.NewUint32(table, "release_id")
	s.AsDefault = field.NewBool(table, "as_default")
	s.Scope = field.NewField(table, "scope")
	s.Namespace = field.NewString(table, "namespace")
	s.Memo = field.NewString(table, "memo")
	s.PublishType = field.NewString(table, "publish_type")
	s.PublishTime = field.NewString(table, "publish_time")
	s.PublishStatus = field.NewString(table, "publish_status")
	s.RejectReason = field.NewString(table, "reject_reason")
	s.Approver = field.NewString(table, "approver")
	s.ApproverProgress = field.NewString(table, "approver_progress")
	s.FinalApprovalTime = field.NewTime(table, "final_approval_time")
	s.ItsmTicketType = field.NewString(table, "itsm_ticket_type")
	s.ItsmTicketUrl = field.NewString(table, "itsm_ticket_url")
	s.ItsmTicketSn = field.NewString(table, "itsm_ticket_sn")
	s.ItsmTicketStatus = field.NewString(table, "itsm_ticket_status")
	s.ItsmTicketStateID = field.NewInt(table, "itsm_ticket_state_id")
	s.PubState = field.NewString(table, "pub_state")
	s.BizID = field.NewUint32(table, "biz_id")
	s.AppID = field.NewUint32(table, "app_id")
	s.StrategySetID = field.NewUint32(table, "strategy_set_id")
	s.Creator = field.NewString(table, "creator")
	s.Reviser = field.NewString(table, "reviser")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *strategy) WithContext(ctx context.Context) IStrategyDo { return s.strategyDo.WithContext(ctx) }

func (s strategy) TableName() string { return s.strategyDo.TableName() }

func (s strategy) Alias() string { return s.strategyDo.Alias() }

func (s strategy) Columns(cols ...field.Expr) gen.Columns { return s.strategyDo.Columns(cols...) }

func (s *strategy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *strategy) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 27)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["release_id"] = s.ReleaseID
	s.fieldMap["as_default"] = s.AsDefault
	s.fieldMap["scope"] = s.Scope
	s.fieldMap["namespace"] = s.Namespace
	s.fieldMap["memo"] = s.Memo
	s.fieldMap["publish_type"] = s.PublishType
	s.fieldMap["publish_time"] = s.PublishTime
	s.fieldMap["publish_status"] = s.PublishStatus
	s.fieldMap["reject_reason"] = s.RejectReason
	s.fieldMap["approver"] = s.Approver
	s.fieldMap["approver_progress"] = s.ApproverProgress
	s.fieldMap["final_approval_time"] = s.FinalApprovalTime
	s.fieldMap["itsm_ticket_type"] = s.ItsmTicketType
	s.fieldMap["itsm_ticket_url"] = s.ItsmTicketUrl
	s.fieldMap["itsm_ticket_sn"] = s.ItsmTicketSn
	s.fieldMap["itsm_ticket_status"] = s.ItsmTicketStatus
	s.fieldMap["itsm_ticket_state_id"] = s.ItsmTicketStateID
	s.fieldMap["pub_state"] = s.PubState
	s.fieldMap["biz_id"] = s.BizID
	s.fieldMap["app_id"] = s.AppID
	s.fieldMap["strategy_set_id"] = s.StrategySetID
	s.fieldMap["creator"] = s.Creator
	s.fieldMap["reviser"] = s.Reviser
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s strategy) clone(db *gorm.DB) strategy {
	s.strategyDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s strategy) replaceDB(db *gorm.DB) strategy {
	s.strategyDo.ReplaceDB(db)
	return s
}

type strategyDo struct{ gen.DO }

type IStrategyDo interface {
	gen.SubQuery
	Debug() IStrategyDo
	WithContext(ctx context.Context) IStrategyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStrategyDo
	WriteDB() IStrategyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStrategyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStrategyDo
	Not(conds ...gen.Condition) IStrategyDo
	Or(conds ...gen.Condition) IStrategyDo
	Select(conds ...field.Expr) IStrategyDo
	Where(conds ...gen.Condition) IStrategyDo
	Order(conds ...field.Expr) IStrategyDo
	Distinct(cols ...field.Expr) IStrategyDo
	Omit(cols ...field.Expr) IStrategyDo
	Join(table schema.Tabler, on ...field.Expr) IStrategyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStrategyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStrategyDo
	Group(cols ...field.Expr) IStrategyDo
	Having(conds ...gen.Condition) IStrategyDo
	Limit(limit int) IStrategyDo
	Offset(offset int) IStrategyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStrategyDo
	Unscoped() IStrategyDo
	Create(values ...*table.Strategy) error
	CreateInBatches(values []*table.Strategy, batchSize int) error
	Save(values ...*table.Strategy) error
	First() (*table.Strategy, error)
	Take() (*table.Strategy, error)
	Last() (*table.Strategy, error)
	Find() ([]*table.Strategy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Strategy, err error)
	FindInBatches(result *[]*table.Strategy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.Strategy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStrategyDo
	Assign(attrs ...field.AssignExpr) IStrategyDo
	Joins(fields ...field.RelationField) IStrategyDo
	Preload(fields ...field.RelationField) IStrategyDo
	FirstOrInit() (*table.Strategy, error)
	FirstOrCreate() (*table.Strategy, error)
	FindByPage(offset int, limit int) (result []*table.Strategy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStrategyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s strategyDo) Debug() IStrategyDo {
	return s.withDO(s.DO.Debug())
}

func (s strategyDo) WithContext(ctx context.Context) IStrategyDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s strategyDo) ReadDB() IStrategyDo {
	return s.Clauses(dbresolver.Read)
}

func (s strategyDo) WriteDB() IStrategyDo {
	return s.Clauses(dbresolver.Write)
}

func (s strategyDo) Session(config *gorm.Session) IStrategyDo {
	return s.withDO(s.DO.Session(config))
}

func (s strategyDo) Clauses(conds ...clause.Expression) IStrategyDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s strategyDo) Returning(value interface{}, columns ...string) IStrategyDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s strategyDo) Not(conds ...gen.Condition) IStrategyDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s strategyDo) Or(conds ...gen.Condition) IStrategyDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s strategyDo) Select(conds ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s strategyDo) Where(conds ...gen.Condition) IStrategyDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s strategyDo) Order(conds ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s strategyDo) Distinct(cols ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s strategyDo) Omit(cols ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s strategyDo) Join(table schema.Tabler, on ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s strategyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s strategyDo) RightJoin(table schema.Tabler, on ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s strategyDo) Group(cols ...field.Expr) IStrategyDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s strategyDo) Having(conds ...gen.Condition) IStrategyDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s strategyDo) Limit(limit int) IStrategyDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s strategyDo) Offset(offset int) IStrategyDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s strategyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStrategyDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s strategyDo) Unscoped() IStrategyDo {
	return s.withDO(s.DO.Unscoped())
}

func (s strategyDo) Create(values ...*table.Strategy) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s strategyDo) CreateInBatches(values []*table.Strategy, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s strategyDo) Save(values ...*table.Strategy) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s strategyDo) First() (*table.Strategy, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.Strategy), nil
	}
}

func (s strategyDo) Take() (*table.Strategy, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.Strategy), nil
	}
}

func (s strategyDo) Last() (*table.Strategy, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.Strategy), nil
	}
}

func (s strategyDo) Find() ([]*table.Strategy, error) {
	result, err := s.DO.Find()
	return result.([]*table.Strategy), err
}

func (s strategyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Strategy, err error) {
	buf := make([]*table.Strategy, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s strategyDo) FindInBatches(result *[]*table.Strategy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s strategyDo) Attrs(attrs ...field.AssignExpr) IStrategyDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s strategyDo) Assign(attrs ...field.AssignExpr) IStrategyDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s strategyDo) Joins(fields ...field.RelationField) IStrategyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s strategyDo) Preload(fields ...field.RelationField) IStrategyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s strategyDo) FirstOrInit() (*table.Strategy, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.Strategy), nil
	}
}

func (s strategyDo) FirstOrCreate() (*table.Strategy, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.Strategy), nil
	}
}

func (s strategyDo) FindByPage(offset int, limit int) (result []*table.Strategy, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s strategyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s strategyDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s strategyDo) Delete(models ...*table.Strategy) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *strategyDo) withDO(do gen.Dao) *strategyDo {
	s.DO = *do.(*gen.DO)
	return s
}
