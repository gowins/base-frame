package hltv

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GgdataCsgoHltvEventModel = (*customGgdataCsgoHltvEventModel)(nil)

type (
	// GgdataCsgoHltvEventModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGgdataCsgoHltvEventModel.
	GgdataCsgoHltvEventModel interface {
		ggdataCsgoHltvEventModel
		FindPageList(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GgdataCsgoHltvEvent, error)
		FindPageListByMap(ctx context.Context, where map[string]interface{}, page, pageSize int64, orderBy string) ([]*GgdataCsgoHltvEvent, error)
	}

	customGgdataCsgoHltvEventModel struct {
		*defaultGgdataCsgoHltvEventModel
	}
)

// NewGgdataCsgoHltvEventModel returns a model for the database table.
func NewGgdataCsgoHltvEventModel(conn sqlx.SqlConn) GgdataCsgoHltvEventModel {
	return &customGgdataCsgoHltvEventModel{
		defaultGgdataCsgoHltvEventModel: newGgdataCsgoHltvEventModel(conn),
	}
}

// FindPageList 根据搜索条件分页查询(SelectBuilder方式)
func (m *defaultGgdataCsgoHltvEventModel) FindPageList(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GgdataCsgoHltvEvent, error) {
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	// 计算offset
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	// 组装sql
	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	// 查询数据
	var resp []*GgdataCsgoHltvEvent
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FindPageListByMap FIndPageListByMap 根据搜索条件分页查询(map的方式)
func (m *defaultGgdataCsgoHltvEventModel) FindPageListByMap(ctx context.Context, where map[string]interface{}, page, pageSize int64, orderBy string) ([]*GgdataCsgoHltvEvent, error) {
	rowBuilder := squirrel.Select("*").From(m.table)
	if len(where) == 0 {
		return m.FindPageList(ctx, rowBuilder, page, pageSize, orderBy)
	}

	// 解析where条件 todo where改造 保持顺序
	for k, v := range where {
		// 断言是否为特殊操作符
		value, ok := v.([]interface{})
		if !ok {
			rowBuilder = rowBuilder.Where(squirrel.Eq{k: v})
			continue
		}

		if op, ok := value[0].(string); ok {
			switch op {
			case ">":
				rowBuilder = rowBuilder.Where(squirrel.Gt{k: value[1]})
			case "<":
				rowBuilder = rowBuilder.Where(squirrel.Lt{k: value[1]})
			case ">=":
				rowBuilder = rowBuilder.Where(squirrel.GtOrEq{k: value[1]})
			case "<=":
				rowBuilder = rowBuilder.Where(squirrel.LtOrEq{k: value[1]})
			default:
				logx.Errorf("不支持的操作符:%s", op)
			}
		}
	}

	return m.FindPageList(ctx, rowBuilder, page, pageSize, orderBy)
}
