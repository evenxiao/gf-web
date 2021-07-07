// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"mypj/app/dao/internal"
)

// adminsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type adminsDao struct {
	*internal.AdminsDao
}

var (
	// Admins is globally public accessible object for table admins operations.
	Admins = &adminsDao{
		internal.Admins,
	}
)

// Fill with you ideas below.