// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
)

// internalSysSettingsDao is internal type for wrapping internal DAO implements.
type internalSysSettingsDao = *internal.SysSettingsDao

// sysSettingsDao is the data access object for table sys_settings.
// You can define custom methods on it to extend its functionality as you wish.
type sysSettingsDao struct {
	internalSysSettingsDao
}

var (
	// SysSettings is globally public accessible object for table sys_settings operations.
	SysSettings = sysSettingsDao{
		internal.NewSysSettingsDao(),
	}
)

// Fill with you ideas below.
