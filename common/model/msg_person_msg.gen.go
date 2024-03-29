// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMsgPersonMsg = "msg_person_msg"

// MsgPersonMsg mapped from table <msg_person_msg>
type MsgPersonMsg struct {
	ID           uint64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" mapstructure:"id"`
	UID          uint32 `gorm:"column:uid;primaryKey;uniqueIndex:uniq_uid_orgin_id_mpm,priority:2;index:idx_uid_create_time_status_mpm,priority:1" json:"uid" mapstructure:"uid"` // 用户id
	FromUID      uint32 `gorm:"column:from_uid;not null" json:"from_uid" mapstructure:"from_uid"`                                                                                 // 发送者用户id
	FromDomain   string `gorm:"column:from_domain" json:"from_domain" mapstructure:"from_domain"`                                                                                 // 用户标识
	FromAvatar   string `gorm:"column:from_avatar;not null" json:"from_avatar" mapstructure:"from_avatar"`                                                                        // 发送者头像
	FromUsername string `gorm:"column:from_username;not null" json:"from_username" mapstructure:"from_username"`                                                                  // 发送者用户名
	FromType     uint8  `gorm:"column:from_type;not null" json:"from_type" mapstructure:"from_type"`                                                                              // 来源类型:0文章 1视频 2赛事 3对阵 4直播 5战绩中心 6社区
	FromID       string `gorm:"column:from_id;not null" json:"from_id" mapstructure:"from_id"`                                                                                    // 来源id
	FromAlias    string `gorm:"column:from_alias;not null" json:"from_alias" mapstructure:"from_alias"`                                                                           // 跳转组装
	FromName     string `gorm:"column:from_name;not null" json:"from_name" mapstructure:"from_name"`                                                                              // 归属名称
	Content      string `gorm:"column:content;not null" json:"content" mapstructure:"content"`                                                                                    // 内容
	Extras       string `gorm:"column:extras" json:"extras" mapstructure:"extras"`                                                                                                // 扩展数据
	Status       uint8  `gorm:"column:status;not null;index:idx_uid_create_time_status_mpm,priority:3" json:"status" mapstructure:"status"`                                       // 0未读 1已读
	Type         uint8  `gorm:"column:type;not null" json:"type" mapstructure:"type"`                                                                                             // 消息类型:0评论 1点赞 2艾特
	CreateTime   uint32 `gorm:"column:create_time;not null;index:idx_uid_create_time_status_mpm,priority:2" json:"create_time" mapstructure:"create_time"`                        // 创建时间
	UpdateTime   uint32 `gorm:"column:update_time;not null;index:idx_update_time_mpm,priority:1" json:"update_time" mapstructure:"update_time"`                                   // 更新时间
	OriginID     string `gorm:"column:origin_id;not null;uniqueIndex:uniq_uid_orgin_id_mpm,priority:1" json:"origin_id" mapstructure:"origin_id"`                                 // 原始id
}

// TableName MsgPersonMsg's table name
func (*MsgPersonMsg) TableName() string {
	return TableNameMsgPersonMsg
}
