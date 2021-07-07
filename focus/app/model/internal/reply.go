// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Reply is the golang structure for table gf_reply.
type Reply struct {
	Id         uint        `orm:"id,primary"  json:"id"`          // 回复ID
	ParentId   uint        `orm:"parent_id"   json:"parent_id"`   // 回复对应的上一级回复ID(没有的话默认为0)
	Title      string      `orm:"title"       json:"title"`       // 回复标题
	Content    string      `orm:"content"     json:"content"`     // 回复内容
	TargetType string      `orm:"target_type" json:"target_type"` // 评论类型: content, reply
	TargetId   uint        `orm:"target_id"   json:"target_id"`   // 对应内容ID，可能回复的是另一个回复，所以这里没有使用content_id
	UserId     uint        `orm:"user_id"     json:"user_id"`     // 网站用户ID
	ZanCount   uint        `orm:"zan_count"   json:"zan_count"`   // 赞
	CaiCount   uint        `orm:"cai_count"   json:"cai_count"`   // 踩
	CreatedAt  *gtime.Time `orm:"created_at"  json:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updated_at"`  //
}
