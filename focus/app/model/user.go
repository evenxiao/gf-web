// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"focus/app/model/internal"
)

const (
	UserStatusOk       = 0 // 用户状态正常
	UserStatusDisabled = 1 // 用户状态禁用
	UserGenderUnknown  = 0 // 性别: 未知
	UserGenderMale     = 1 // 性别: 男
	UserGenderFemale   = 2 // 性别: 女
)

// User is the golang structure for table gf_user.
type User internal.User