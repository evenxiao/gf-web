// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// RoleMenu is the golang structure for table role_menu.
type RoleMenu struct {
    Id         uint        `orm:"id,primary"  json:"id"`          //                 
    RoleId     uint        `orm:"role_id"     json:"role_id"`     //                 
    MenuId     uint        `orm:"menu_id"     json:"menu_id"`     //                 
    CreateTime *gtime.Time `orm:"create_time" json:"create_time"` //                 
    Status     uint        `orm:"status"      json:"status"`      // 1 有效  0 无效  
}