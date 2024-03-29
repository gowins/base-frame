// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMsgPersonCount = "msg_person_count"

// MsgPersonCount mapped from table <msg_person_count>
type MsgPersonCount struct {
	ID         uint64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" mapstructure:"id"`
	UID        uint32 `gorm:"column:uid;primaryKey;uniqueIndex:uniq_uid,priority:1" json:"uid" mapstructure:"uid"`                        // 用户id
	Sum        uint32 `gorm:"column:sum;not null" json:"sum" mapstructure:"sum"`                                                          // 消息总数
	Unread     uint32 `gorm:"column:unread;not null" json:"unread" mapstructure:"unread"`                                                 // 未读数量
	CreateTime uint32 `gorm:"column:create_time;not null" json:"create_time" mapstructure:"create_time"`                                  // 创建时间
	UpdateTime uint32 `gorm:"column:update_time;not null;index:idx_update_time,priority:1" json:"update_time" mapstructure:"update_time"` // 更新时间
}

// TableName MsgPersonCount's table name
func (*MsgPersonCount) TableName() string {
	return TableNameMsgPersonCount
}
