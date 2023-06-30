// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
)

// internalSysPersonLicenseDao is internal type for wrapping internal DAO implements.
type internalSysPersonLicenseDao = *internal.SysPersonLicenseDao

// sysPersonLicenseDao is the data access object for table sys_person_license.
// You can define custom methods on it to extend its functionality as you wish.
type sysPersonLicenseDao struct {
	internalSysPersonLicenseDao
}

var (
	// SysPersonLicense is globally public accessible object for table sys_person_license operations.
	SysPersonLicense = sysPersonLicenseDao{
		internal.NewSysPersonLicenseDao(),
	}
)

// Fill with you ideas below.
