// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

const (
	SessionNoticeTypeSuccess = "success"
	SessionNoticeTypeInfo    = "primary"
	SessionNoticeTypeWarn    = "warning"
	SessionNoticeTypeError   = "danger"
)

// 存放在Session中的提示信息，往往使用后则删除
type SessionNotice struct {
	Type    string // 消息类型
	Content string // 消息内容
}