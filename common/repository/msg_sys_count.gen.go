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

func newMsgSysCount(db *gorm.DB, opts ...gen.DOOption) msgSysCount {
	_msgSysCount := msgSysCount{}

	_msgSysCount.msgSysCountDo.UseDB(db, opts...)
	_msgSysCount.msgSysCountDo.UseModel(&model.MsgSysCount{})

	tableName := _msgSysCount.msgSysCountDo.TableName()
	_msgSysCount.ALL = field.NewAsterisk(tableName)
	_msgSysCount.ID = field.NewUint64(tableName, "id")
	_msgSysCount.UID = field.NewUint32(tableName, "uid")
	_msgSysCount.Sum = field.NewUint32(tableName, "sum")
	_msgSysCount.Unread = field.NewUint32(tableName, "unread")
	_msgSysCount.CreateTime = field.NewUint32(tableName, "create_time")
	_msgSysCount.UpdateTime = field.NewUint32(tableName, "update_time")

	_msgSysCount.fillFieldMap()

	return _msgSysCount
}

type msgSysCount struct {
	msgSysCountDo msgSysCountDo

	ALL        field.Asterisk
	ID         field.Uint64
	UID        field.Uint32 // 用户id
	Sum        field.Uint32 // 消息总数
	Unread     field.Uint32 // 未读数量
	CreateTime field.Uint32 // 创建时间
	UpdateTime field.Uint32 // 更新时间

	fieldMap map[string]field.Expr
}

func (m msgSysCount) Table(newTableName string) *msgSysCount {
	m.msgSysCountDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msgSysCount) As(alias string) *msgSysCount {
	m.msgSysCountDo.DO = *(m.msgSysCountDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msgSysCount) updateTableName(table string) *msgSysCount {
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

func (m *msgSysCount) WithContext(ctx context.Context) *msgSysCountDo {
	return m.msgSysCountDo.WithContext(ctx)
}

func (m msgSysCount) TableName() string { return m.msgSysCountDo.TableName() }

func (m msgSysCount) Alias() string { return m.msgSysCountDo.Alias() }

func (m *msgSysCount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msgSysCount) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["uid"] = m.UID
	m.fieldMap["sum"] = m.Sum
	m.fieldMap["unread"] = m.Unread
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
}

func (m msgSysCount) clone(db *gorm.DB) msgSysCount {
	m.msgSysCountDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msgSysCount) replaceDB(db *gorm.DB) msgSysCount {
	m.msgSysCountDo.ReplaceDB(db)
	return m
}

type msgSysCountDo struct{ gen.DO }

func (m msgSysCountDo) Debug() *msgSysCountDo {
	return m.withDO(m.DO.Debug())
}

func (m msgSysCountDo) WithContext(ctx context.Context) *msgSysCountDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgSysCountDo) ReadDB() *msgSysCountDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgSysCountDo) WriteDB() *msgSysCountDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgSysCountDo) Session(config *gorm.Session) *msgSysCountDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgSysCountDo) Clauses(conds ...clause.Expression) *msgSysCountDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgSysCountDo) Returning(value interface{}, columns ...string) *msgSysCountDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgSysCountDo) Not(conds ...gen.Condition) *msgSysCountDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgSysCountDo) Or(conds ...gen.Condition) *msgSysCountDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgSysCountDo) Select(conds ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgSysCountDo) Where(conds ...gen.Condition) *msgSysCountDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgSysCountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *msgSysCountDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m msgSysCountDo) Order(conds ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgSysCountDo) Distinct(cols ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgSysCountDo) Omit(cols ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgSysCountDo) Join(table schema.Tabler, on ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgSysCountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgSysCountDo) RightJoin(table schema.Tabler, on ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgSysCountDo) Group(cols ...field.Expr) *msgSysCountDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgSysCountDo) Having(conds ...gen.Condition) *msgSysCountDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgSysCountDo) Limit(limit int) *msgSysCountDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgSysCountDo) Offset(offset int) *msgSysCountDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgSysCountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *msgSysCountDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgSysCountDo) Unscoped() *msgSysCountDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgSysCountDo) Create(values ...*model.MsgSysCount) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgSysCountDo) CreateInBatches(values []*model.MsgSysCount, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgSysCountDo) Save(values ...*model.MsgSysCount) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgSysCountDo) First() (*model.MsgSysCount, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSysCount), nil
	}
}

func (m msgSysCountDo) Take() (*model.MsgSysCount, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSysCount), nil
	}
}

func (m msgSysCountDo) Last() (*model.MsgSysCount, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSysCount), nil
	}
}

func (m msgSysCountDo) Find() ([]*model.MsgSysCount, error) {
	result, err := m.DO.Find()
	return result.([]*model.MsgSysCount), err
}

func (m msgSysCountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MsgSysCount, err error) {
	buf := make([]*model.MsgSysCount, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgSysCountDo) FindInBatches(result *[]*model.MsgSysCount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgSysCountDo) Attrs(attrs ...field.AssignExpr) *msgSysCountDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgSysCountDo) Assign(attrs ...field.AssignExpr) *msgSysCountDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgSysCountDo) Joins(fields ...field.RelationField) *msgSysCountDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgSysCountDo) Preload(fields ...field.RelationField) *msgSysCountDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgSysCountDo) FirstOrInit() (*model.MsgSysCount, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSysCount), nil
	}
}

func (m msgSysCountDo) FirstOrCreate() (*model.MsgSysCount, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSysCount), nil
	}
}

func (m msgSysCountDo) FindByPage(offset int, limit int) (result []*model.MsgSysCount, count int64, err error) {
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

func (m msgSysCountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgSysCountDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgSysCountDo) Delete(models ...*model.MsgSysCount) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgSysCountDo) withDO(do gen.Dao) *msgSysCountDo {
	m.DO = *do.(*gen.DO)
	return m
}
