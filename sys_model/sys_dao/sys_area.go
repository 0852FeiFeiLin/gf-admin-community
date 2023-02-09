// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
)

// internalSysAreaDao is internal type for wrapping internal DAO implements.
type internalSysAreaDao = *internal.SysAreaDao

// sysAreaDao is the data access object for table sys_area.
// You can define custom methods on it to extend its functionality as you wish.
type sysAreaDao struct {
	internalSysAreaDao
}

var (
	// SysArea is globally public accessible object for table sys_area operations.
	SysArea = sysAreaDao{
		internal.NewSysAreaDao(),
	}
)

// Fill with you ideas below.
