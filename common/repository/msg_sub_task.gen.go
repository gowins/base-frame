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

func newMsgSubTask(db *gorm.DB, opts ...gen.DOOption) msgSubTask {
	_msgSubTask := msgSubTask{}

	_msgSubTask.msgSubTaskDo.UseDB(db, opts...)
	_msgSubTask.msgSubTaskDo.UseModel(&model.MsgSubTask{})

	tableName := _msgSubTask.msgSubTaskDo.TableName()
	_msgSubTask.ALL = field.NewAsterisk(tableName)
	_msgSubTask.ID = field.NewUint32(tableName, "id")
	_msgSubTask.TaskID = field.NewUint32(tableName, "task_id")
	_msgSubTask.TemplateID = field.NewUint32(tableName, "template_id")
	_msgSubTask.Status = field.NewInt8(tableName, "status")
	_msgSubTask.Creator = field.NewString(tableName, "creator")
	_msgSubTask.CreatedAt = field.NewTime(tableName, "created_at")
	_msgSubTask.UpdatedAt = field.NewTime(tableName, "updated_at")
	_msgSubTask.DeletedAt = field.NewField(tableName, "deleted_at")

	_msgSubTask.fillFieldMap()

	return _msgSubTask
}

type msgSubTask struct {
	msgSubTaskDo msgSubTaskDo

	ALL        field.Asterisk
	ID         field.Uint32
	TaskID     field.Uint32 // 任务id
	TemplateID field.Uint32 // 模版id
	Status     field.Int8   // 任务状态： 0-启用 1-停用
	Creator    field.String // 创建人
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 更新时间
	DeletedAt  field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (m msgSubTask) Table(newTableName string) *msgSubTask {
	m.msgSubTaskDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msgSubTask) As(alias string) *msgSubTask {
	m.msgSubTaskDo.DO = *(m.msgSubTaskDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msgSubTask) updateTableName(table string) *msgSubTask {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint32(table, "id")
	m.TaskID = field.NewUint32(table, "task_id")
	m.TemplateID = field.NewUint32(table, "template_id")
	m.Status = field.NewInt8(table, "status")
	m.Creator = field.NewString(table, "creator")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *msgSubTask) WithContext(ctx context.Context) *msgSubTaskDo {
	return m.msgSubTaskDo.WithContext(ctx)
}

func (m msgSubTask) TableName() string { return m.msgSubTaskDo.TableName() }

func (m msgSubTask) Alias() string { return m.msgSubTaskDo.Alias() }

func (m *msgSubTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msgSubTask) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 8)
	m.fieldMap["id"] = m.ID
	m.fieldMap["task_id"] = m.TaskID
	m.fieldMap["template_id"] = m.TemplateID
	m.fieldMap["status"] = m.Status
	m.fieldMap["creator"] = m.Creator
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
}

func (m msgSubTask) clone(db *gorm.DB) msgSubTask {
	m.msgSubTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msgSubTask) replaceDB(db *gorm.DB) msgSubTask {
	m.msgSubTaskDo.ReplaceDB(db)
	return m
}

type msgSubTaskDo struct{ gen.DO }

func (m msgSubTaskDo) Debug() *msgSubTaskDo {
	return m.withDO(m.DO.Debug())
}

func (m msgSubTaskDo) WithContext(ctx context.Context) *msgSubTaskDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgSubTaskDo) ReadDB() *msgSubTaskDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgSubTaskDo) WriteDB() *msgSubTaskDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgSubTaskDo) Session(config *gorm.Session) *msgSubTaskDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgSubTaskDo) Clauses(conds ...clause.Expression) *msgSubTaskDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgSubTaskDo) Returning(value interface{}, columns ...string) *msgSubTaskDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgSubTaskDo) Not(conds ...gen.Condition) *msgSubTaskDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgSubTaskDo) Or(conds ...gen.Condition) *msgSubTaskDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgSubTaskDo) Select(conds ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgSubTaskDo) Where(conds ...gen.Condition) *msgSubTaskDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgSubTaskDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *msgSubTaskDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m msgSubTaskDo) Order(conds ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgSubTaskDo) Distinct(cols ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgSubTaskDo) Omit(cols ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgSubTaskDo) Join(table schema.Tabler, on ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgSubTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgSubTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgSubTaskDo) Group(cols ...field.Expr) *msgSubTaskDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgSubTaskDo) Having(conds ...gen.Condition) *msgSubTaskDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgSubTaskDo) Limit(limit int) *msgSubTaskDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgSubTaskDo) Offset(offset int) *msgSubTaskDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgSubTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *msgSubTaskDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgSubTaskDo) Unscoped() *msgSubTaskDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgSubTaskDo) Create(values ...*model.MsgSubTask) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgSubTaskDo) CreateInBatches(values []*model.MsgSubTask, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgSubTaskDo) Save(values ...*model.MsgSubTask) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgSubTaskDo) First() (*model.MsgSubTask, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSubTask), nil
	}
}

func (m msgSubTaskDo) Take() (*model.MsgSubTask, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSubTask), nil
	}
}

func (m msgSubTaskDo) Last() (*model.MsgSubTask, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSubTask), nil
	}
}

func (m msgSubTaskDo) Find() ([]*model.MsgSubTask, error) {
	result, err := m.DO.Find()
	return result.([]*model.MsgSubTask), err
}

func (m msgSubTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MsgSubTask, err error) {
	buf := make([]*model.MsgSubTask, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgSubTaskDo) FindInBatches(result *[]*model.MsgSubTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgSubTaskDo) Attrs(attrs ...field.AssignExpr) *msgSubTaskDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgSubTaskDo) Assign(attrs ...field.AssignExpr) *msgSubTaskDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgSubTaskDo) Joins(fields ...field.RelationField) *msgSubTaskDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgSubTaskDo) Preload(fields ...field.RelationField) *msgSubTaskDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgSubTaskDo) FirstOrInit() (*model.MsgSubTask, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSubTask), nil
	}
}

func (m msgSubTaskDo) FirstOrCreate() (*model.MsgSubTask, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgSubTask), nil
	}
}

func (m msgSubTaskDo) FindByPage(offset int, limit int) (result []*model.MsgSubTask, count int64, err error) {
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

func (m msgSubTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgSubTaskDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgSubTaskDo) Delete(models ...*model.MsgSubTask) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgSubTaskDo) withDO(do gen.Dao) *msgSubTaskDo {
	m.DO = *do.(*gen.DO)
	return m
}
