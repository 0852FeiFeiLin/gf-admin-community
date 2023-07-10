// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
)

type SysOrganizationDao = dao_interface.TIDao[internal.SysOrganizationColumns]

func NewSysOrganization(dao ...dao_interface.IDao) SysOrganizationDao {
	return (SysOrganizationDao)(internal.NewSysOrganizationDao(dao...))
}

var (
	SysOrganization = NewSysOrganization()
)
