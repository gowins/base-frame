// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"base-frame/common/model"
)

func newMsgPersonCount(db *gorm.DB, opts ...gen.DOOption) msgPersonCount {
	_msgPersonCount := msgPersonCount{}

	_msgPersonCount.msgPersonCountDo.UseDB(db, opts...)
	_msgPersonCount.msgPersonCountDo.UseModel(&model.MsgPersonCount{})

	tableName := _msgPersonCount.msgPersonCountDo.TableName()
	_msgPersonCount.ALL = field.NewAsterisk(tableName)
	_msgPersonCount.ID = field.NewUint64(tableName, "id")
	_msgPersonCount.UID = field.NewUint32(tableName, "uid")
	_msgPersonCount.Sum = field.NewUint32(tableName, "sum")
	_msgPersonCount.Unread = field.NewUint32(tableName, "unread")
	_msgPersonCount.CreateTime = field.NewUint32(tableName, "create_time")
	_msgPersonCount.UpdateTime = field.NewUint32(tableName, "update_time")

	_msgPersonCount.fillFieldMap()

	return _msgPersonCount
}

type msgPersonCount struct {
	msgPersonCountDo msgPersonCountDo

	ALL        field.Asterisk
	ID         field.Uint64
	UID        field.Uint32 // 用户id
	Sum        field.Uint32 // 消息总数
	Unread     field.Uint32 // 未读数量
	CreateTime field.Uint32 // 创建时间
	UpdateTime field.Uint32 // 更新时间

	fieldMap map[string]field.Expr
}

func (m msgPersonCount) Table(newTableName string) *msgPersonCount {
	m.msgPersonCountDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msgPersonCount) As(alias string) *msgPersonCount {
	m.msgPersonCountDo.DO = *(m.msgPersonCountDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msgPersonCount) updateTableName(table string) *msgPersonCount {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint64(table, "id")
	m.UID = field.NewUint32(table, "uid")
	m.Sum = field.NewUint32(table, "sum")
	m.Unread = field.NewUint32(table, "unread")
	m.CreateTime = field.NewUint32(table, "create_time")
	m.UpdateTime = field.NewUint32(table, "update_time")

	m.fillFieldMap()

	return m
}

func (m *msgPersonCount) WithContext(ctx context.Context) *msgPersonCountDo {
	return m.msgPersonCountDo.WithContext(ctx)
}

func (m msgPersonCount) TableName() string { return m.msgPersonCountDo.TableName() }

func (m msgPersonCount) Alias() string { return m.msgPersonCountDo.Alias() }

func (m *msgPersonCount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msgPersonCount) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["uid"] = m.UID
	m.fieldMap["sum"] = m.Sum
	m.fieldMap["unread"] = m.Unread
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
}

func (m msgPersonCount) clone(db *gorm.DB) msgPersonCount {
	m.msgPersonCountDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msgPersonCount) replaceDB(db *gorm.DB) msgPersonCount {
	m.msgPersonCountDo.ReplaceDB(db)
	return m
}

type msgPersonCountDo struct{ gen.DO }

func (m msgPersonCountDo) Debug() *msgPersonCountDo {
	return m.withDO(m.DO.Debug())
}

func (m msgPersonCountDo) WithContext(ctx context.Context) *msgPersonCountDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgPersonCountDo) ReadDB() *msgPersonCountDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgPersonCountDo) WriteDB() *msgPersonCountDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgPersonCountDo) Session(config *gorm.Session) *msgPersonCountDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgPersonCountDo) Clauses(conds ...clause.Expression) *msgPersonCountDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgPersonCountDo) Returning(value interface{}, columns ...string) *msgPersonCountDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgPersonCountDo) Not(conds ...gen.Condition) *msgPersonCountDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgPersonCountDo) Or(conds ...gen.Condition) *msgPersonCountDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgPersonCountDo) Select(conds ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgPersonCountDo) Where(conds ...gen.Condition) *msgPersonCountDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgPersonCountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *msgPersonCountDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m msgPersonCountDo) Order(conds ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgPersonCountDo) Distinct(cols ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgPersonCountDo) Omit(cols ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgPersonCountDo) Join(table schema.Tabler, on ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgPersonCountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgPersonCountDo) RightJoin(table schema.Tabler, on ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgPersonCountDo) Group(cols ...field.Expr) *msgPersonCountDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgPersonCountDo) Having(conds ...gen.Condition) *msgPersonCountDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgPersonCountDo) Limit(limit int) *msgPersonCountDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgPersonCountDo) Offset(offset int) *msgPersonCountDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgPersonCountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *msgPersonCountDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgPersonCountDo) Unscoped() *msgPersonCountDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgPersonCountDo) Create(values ...*model.MsgPersonCount) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgPersonCountDo) CreateInBatches(values []*model.MsgPersonCount, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgPersonCountDo) Save(values ...*model.MsgPersonCount) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgPersonCountDo) First() (*model.MsgPersonCount, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgPersonCount), nil
	}
}

func (m msgPersonCountDo) Take() (*model.MsgPersonCount, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgPersonCount), nil
	}
}

func (m msgPersonCountDo) Last() (*model.MsgPersonCount, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgPersonCount), nil
	}
}

func (m msgPersonCountDo) Find() ([]*model.MsgPersonCount, error) {
	result, err := m.DO.Find()
	return result.([]*model.MsgPersonCount), err
}

func (m msgPersonCountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MsgPersonCount, err error) {
	buf := make([]*model.MsgPersonCount, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgPersonCountDo) FindInBatches(result *[]*model.MsgPersonCount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgPersonCountDo) Attrs(attrs ...field.AssignExpr) *msgPersonCountDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgPersonCountDo) Assign(attrs ...field.AssignExpr) *msgPersonCountDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgPersonCountDo) Joins(fields ...field.RelationField) *msgPersonCountDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgPersonCountDo) Preload(fields ...field.RelationField) *msgPersonCountDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgPersonCountDo) FirstOrInit() (*model.MsgPersonCount, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgPersonCount), nil
	}
}

func (m msgPersonCountDo) FirstOrCreate() (*model.MsgPersonCount, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgPersonCount), nil
	}
}

func (m msgPersonCountDo) FindByPage(offset int, limit int) (result []*model.MsgPersonCount, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m msgPersonCountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgPersonCountDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgPersonCountDo) Delete(models ...*model.MsgPersonCount) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgPersonCountDo) withDO(do gen.Dao) *msgPersonCountDo {
	m.DO = *do.(*gen.DO)
	return m
}
