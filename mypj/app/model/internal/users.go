// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
    UserId         uint        `orm:"user_id,primary"       json:"user_id"`          //
    UserName   string      `orm:"user_name,unique" json:"user_name"`   // 用户名    
    Password   string      `orm:"password"         json:"password"`    // 密码      
    Email      string      `orm:"email"            json:"email"`       // 邮箱      
    Mobile     string      `orm:"mobile"           json:"mobile"`      // 手机号    
    RegTime    *gtime.Time `orm:"reg_time"         json:"reg_time"`    // 注册时间  
    CreateTime *gtime.Time `orm:"create_time"      json:"create_time"` // 创建时间  
    UpdateTime *gtime.Time `orm:"update_time"      json:"update_time"` // 更新时间  
}