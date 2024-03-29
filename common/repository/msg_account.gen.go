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

func newMsgAccount(db *gorm.DB, opts ...gen.DOOption) msgAccount {
	_msgAccount := msgAccount{}

	_msgAccount.msgAccountDo.UseDB(db, opts...)
	_msgAccount.msgAccountDo.UseModel(&model.MsgAccount{})

	tableName := _msgAccount.msgAccountDo.TableName()
	_msgAccount.ALL = field.NewAsterisk(tableName)
	_msgAccount.ID = field.NewUint32(tableName, "id")
	_msgAccount.Tag = field.NewString(tableName, "tag")
	_msgAccount.Name = field.NewString(tableName, "name")
	_msgAccount.Remark = field.NewString(tableName, "remark")
	_msgAccount.Vendor = field.NewString(tableName, "vendor")
	_msgAccount.Config = field.NewString(tableName, "config")
	_msgAccount.Status = field.NewInt8(tableName, "status")
	_msgAccount.Creator = field.NewString(tableName, "creator")
	_msgAccount.CreatedAt = field.NewTime(tableName, "created_at")
	_msgAccount.UpdatedAt = field.NewTime(tableName, "updated_at")
	_msgAccount.DeletedAt = field.NewField(tableName, "deleted_at")

	_msgAccount.fillFieldMap()

	return _msgAccount
}

type msgAccount struct {
	msgAccountDo msgAccountDo

	ALL       field.Asterisk
	ID        field.Uint32
	Tag       field.String // 账号标识
	Name      field.String // 账号名称
	Remark    field.String // 账号备注
	Vendor    field.String // 厂商
	Config    field.String // 账号配置
	Status    field.Int8   // 任务状态： 0-启用 1-停用
	Creator   field.String // 创建人
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (m msgAccount) Table(newTableName string) *msgAccount {
	m.msgAccountDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msgAccount) As(alias string) *msgAccount {
	m.msgAccountDo.DO = *(m.msgAccountDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msgAccount) updateTableName(table string) *msgAccount {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint32(table, "id")
	m.Tag = field.NewString(table, "tag")
	m.Name = field.NewString(table, "name")
	m.Remark = field.NewString(table, "remark")
	m.Vendor = field.NewString(table, "vendor")
	m.Config = field.NewString(table, "config")
	m.Status = field.NewInt8(table, "status")
	m.Creator = field.NewString(table, "creator")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *msgAccount) WithContext(ctx context.Context) *msgAccountDo {
	return m.msgAccountDo.WithContext(ctx)
}

func (m msgAccount) TableName() string { return m.msgAccountDo.TableName() }

func (m msgAccount) Alias() string { return m.msgAccountDo.Alias() }

func (m *msgAccount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msgAccount) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 11)
	m.fieldMap["id"] = m.ID
	m.fieldMap["tag"] = m.Tag
	m.fieldMap["name"] = m.Name
	m.fieldMap["remark"] = m.Remark
	m.fieldMap["vendor"] = m.Vendor
	m.fieldMap["config"] = m.Config
	m.fieldMap["status"] = m.Status
	m.fieldMap["creator"] = m.Creator
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
}

func (m msgAccount) clone(db *gorm.DB) msgAccount {
	m.msgAccountDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msgAccount) replaceDB(db *gorm.DB) msgAccount {
	m.msgAccountDo.ReplaceDB(db)
	return m
}

type msgAccountDo struct{ gen.DO }

func (m msgAccountDo) Debug() *msgAccountDo {
	return m.withDO(m.DO.Debug())
}

func (m msgAccountDo) WithContext(ctx context.Context) *msgAccountDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgAccountDo) ReadDB() *msgAccountDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgAccountDo) WriteDB() *msgAccountDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgAccountDo) Session(config *gorm.Session) *msgAccountDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgAccountDo) Clauses(conds ...clause.Expression) *msgAccountDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgAccountDo) Returning(value interface{}, columns ...string) *msgAccountDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgAccountDo) Not(conds ...gen.Condition) *msgAccountDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgAccountDo) Or(conds ...gen.Condition) *msgAccountDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgAccountDo) Select(conds ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgAccountDo) Where(conds ...gen.Condition) *msgAccountDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgAccountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *msgAccountDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m msgAccountDo) Order(conds ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgAccountDo) Distinct(cols ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgAccountDo) Omit(cols ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgAccountDo) Join(table schema.Tabler, on ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgAccountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgAccountDo) RightJoin(table schema.Tabler, on ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgAccountDo) Group(cols ...field.Expr) *msgAccountDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgAccountDo) Having(conds ...gen.Condition) *msgAccountDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgAccountDo) Limit(limit int) *msgAccountDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgAccountDo) Offset(offset int) *msgAccountDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgAccountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *msgAccountDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgAccountDo) Unscoped() *msgAccountDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgAccountDo) Create(values ...*model.MsgAccount) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgAccountDo) CreateInBatches(values []*model.MsgAccount, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgAccountDo) Save(values ...*model.MsgAccount) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgAccountDo) First() (*model.MsgAccount, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgAccount), nil
	}
}

func (m msgAccountDo) Take() (*model.MsgAccount, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgAccount), nil
	}
}

func (m msgAccountDo) Last() (*model.MsgAccount, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgAccount), nil
	}
}

func (m msgAccountDo) Find() ([]*model.MsgAccount, error) {
	result, err := m.DO.Find()
	return result.([]*model.MsgAccount), err
}

func (m msgAccountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MsgAccount, err error) {
	buf := make([]*model.MsgAccount, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgAccountDo) FindInBatches(result *[]*model.MsgAccount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgAccountDo) Attrs(attrs ...field.AssignExpr) *msgAccountDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgAccountDo) Assign(attrs ...field.AssignExpr) *msgAccountDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgAccountDo) Joins(fields ...field.RelationField) *msgAccountDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgAccountDo) Preload(fields ...field.RelationField) *msgAccountDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgAccountDo) FirstOrInit() (*model.MsgAccount, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgAccount), nil
	}
}

func (m msgAccountDo) FirstOrCreate() (*model.MsgAccount, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgAccount), nil
	}
}

func (m msgAccountDo) FindByPage(offset int, limit int) (result []*model.MsgAccount, count int64, err error) {
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

func (m msgAccountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgAccountDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgAccountDo) Delete(models ...*model.MsgAccount) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgAccountDo) withDO(do gen.Dao) *msgAccountDo {
	m.DO = *do.(*gen.DO)
	return m
}
