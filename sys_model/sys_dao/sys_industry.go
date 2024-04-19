// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
)

type SysIndustryDao = dao_interface.TIDao[internal.SysIndustryColumns]

func NewSysIndustry(dao ...dao_interface.IDao) SysIndustryDao {
	return (SysIndustryDao)(internal.NewSysIndustryDao(dao...))
}

var (
	SysIndustry = NewSysIndustry()
)
