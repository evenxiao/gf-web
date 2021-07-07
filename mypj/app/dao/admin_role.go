// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"mypj/app/dao/internal"
)

// adminRoleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type adminRoleDao struct {
	*internal.AdminRoleDao
}

var (
	// AdminRole is globally public accessible object for table admin_role operations.
	AdminRole = &adminRoleDao{
		internal.AdminRole,
	}
)

// Fill with you ideas below.