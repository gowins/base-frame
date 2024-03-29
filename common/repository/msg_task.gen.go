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

func newMsgTask(db *gorm.DB, opts ...gen.DOOption) msgTask {
	_msgTask := msgTask{}

	_msgTask.msgTaskDo.UseDB(db, opts...)
	_msgTask.msgTaskDo.UseModel(&model.MsgTask{})

	tableName := _msgTask.msgTaskDo.TableName()
	_msgTask.ALL = field.NewAsterisk(tableName)
	_msgTask.ID = field.NewUint32(tableName, "id")
	_msgTask.Name = field.NewString(tableName, "name")
	_msgTask.Code = field.NewString(tableName, "code")
	_msgTask.Params = field.NewString(tableName, "params")
	_msgTask.Type = field.NewInt8(tableName, "type")
	_msgTask.Status = field.NewInt8(tableName, "status")
	_msgTask.Creator = field.NewString(tableName, "creator")
	_msgTask.CreatedAt = field.NewTime(tableName, "created_at")
	_msgTask.UpdatedAt = field.NewTime(tableName, "updated_at")
	_msgTask.DeletedAt = field.NewField(tableName, "deleted_at")

	_msgTask.fillFieldMap()

	return _msgTask
}

type msgTask struct {
	msgTaskDo msgTaskDo

	ALL       field.Asterisk
	ID        field.Uint32
	Name      field.String // 任务名称
	Code      field.String // 任务码
	Params    field.String // 任务参数
	Type      field.Int8   // 任务类型：0-系统发送 1-人工发送
	Status    field.Int8   // 任务状态： 0-启用 1-停用
	Creator   field.String // 创建人
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (m msgTask) Table(newTableName string) *msgTask {
	m.msgTaskDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msgTask) As(alias string) *msgTask {
	m.msgTaskDo.DO = *(m.msgTaskDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msgTask) updateTableName(table string) *msgTask {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint32(table, "id")
	m.Name = field.NewString(table, "name")
	m.Code = field.NewString(table, "code")
	m.Params = field.NewString(table, "params")
	m.Type = field.NewInt8(table, "type")
	m.Status = field.NewInt8(table, "status")
	m.Creator = field.NewString(table, "creator")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *msgTask) WithContext(ctx context.Context) *msgTaskDo { return m.msgTaskDo.WithContext(ctx) }

func (m msgTask) TableName() string { return m.msgTaskDo.TableName() }

func (m msgTask) Alias() string { return m.msgTaskDo.Alias() }

func (m *msgTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msgTask) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 10)
	m.fieldMap["id"] = m.ID
	m.fieldMap["name"] = m.Name
	m.fieldMap["code"] = m.Code
	m.fieldMap["params"] = m.Params
	m.fieldMap["type"] = m.Type
	m.fieldMap["status"] = m.Status
	m.fieldMap["creator"] = m.Creator
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
}

func (m msgTask) clone(db *gorm.DB) msgTask {
	m.msgTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msgTask) replaceDB(db *gorm.DB) msgTask {
	m.msgTaskDo.ReplaceDB(db)
	return m
}

type msgTaskDo struct{ gen.DO }

func (m msgTaskDo) Debug() *msgTaskDo {
	return m.withDO(m.DO.Debug())
}

func (m msgTaskDo) WithContext(ctx context.Context) *msgTaskDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgTaskDo) ReadDB() *msgTaskDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgTaskDo) WriteDB() *msgTaskDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgTaskDo) Session(config *gorm.Session) *msgTaskDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgTaskDo) Clauses(conds ...clause.Expression) *msgTaskDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgTaskDo) Returning(value interface{}, columns ...string) *msgTaskDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgTaskDo) Not(conds ...gen.Condition) *msgTaskDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgTaskDo) Or(conds ...gen.Condition) *msgTaskDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgTaskDo) Select(conds ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgTaskDo) Where(conds ...gen.Condition) *msgTaskDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgTaskDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *msgTaskDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m msgTaskDo) Order(conds ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgTaskDo) Distinct(cols ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgTaskDo) Omit(cols ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgTaskDo) Join(table schema.Tabler, on ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgTaskDo) Group(cols ...field.Expr) *msgTaskDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgTaskDo) Having(conds ...gen.Condition) *msgTaskDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgTaskDo) Limit(limit int) *msgTaskDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgTaskDo) Offset(offset int) *msgTaskDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *msgTaskDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgTaskDo) Unscoped() *msgTaskDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgTaskDo) Create(values ...*model.MsgTask) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgTaskDo) CreateInBatches(values []*model.MsgTask, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgTaskDo) Save(values ...*model.MsgTask) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgTaskDo) First() (*model.MsgTask, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgTask), nil
	}
}

func (m msgTaskDo) Take() (*model.MsgTask, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgTask), nil
	}
}

func (m msgTaskDo) Last() (*model.MsgTask, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgTask), nil
	}
}

func (m msgTaskDo) Find() ([]*model.MsgTask, error) {
	result, err := m.DO.Find()
	return result.([]*model.MsgTask), err
}

func (m msgTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MsgTask, err error) {
	buf := make([]*model.MsgTask, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgTaskDo) FindInBatches(result *[]*model.MsgTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgTaskDo) Attrs(attrs ...field.AssignExpr) *msgTaskDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgTaskDo) Assign(attrs ...field.AssignExpr) *msgTaskDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgTaskDo) Joins(fields ...field.RelationField) *msgTaskDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgTaskDo) Preload(fields ...field.RelationField) *msgTaskDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgTaskDo) FirstOrInit() (*model.MsgTask, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgTask), nil
	}
}

func (m msgTaskDo) FirstOrCreate() (*model.MsgTask, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgTask), nil
	}
}

func (m msgTaskDo) FindByPage(offset int, limit int) (result []*model.MsgTask, count int64, err error) {
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

func (m msgTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgTaskDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgTaskDo) Delete(models ...*model.MsgTask) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgTaskDo) withDO(do gen.Dao) *msgTaskDo {
	m.DO = *do.(*gen.DO)
	return m
}
