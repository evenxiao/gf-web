// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal



// AdminRole is the golang structure for table admin_role.
type AdminRole struct {
    Id      uint `orm:"id,primary" json:"id"`       //                
    AdminId uint `orm:"admin_id"   json:"admin_id"` //                
    RoleId  uint `orm:"role_id"    json:"role_id"`  //                
    Status  uint `orm:"status"     json:"status"`   // 1 有效 0 无效  
}