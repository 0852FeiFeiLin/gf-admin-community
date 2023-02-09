// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao is the data access object for table sys_role.
type SysRoleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleColumns defines and stores column names for table sys_role.
type SysRoleColumns struct {
	Id          string //
	Name        string // 名称
	Description string // 描述
	IsSystem    string // 是否默认角色，true仅能修改名称，不允许删除和修改
	UpdatedAt   string //
	CreatedAt   string //
	UnionMainId string // 主体id
}

// sysRoleColumns holds the columns for table sys_role.
var sysRoleColumns = SysRoleColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	IsSystem:    "is_system",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
	UnionMainId: "union_main_id",
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao(proxy ...dao_interface.IDao) *SysRoleDao {
	var dao *SysRoleDao
	if proxy != nil {
		dao = &SysRoleDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: sysRoleColumns,
		}
		return dao
	}

	return &SysRoleDao{
		group:   "default",
		table:   "sys_role",
		columns: sysRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRoleDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRoleDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRoleDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	model := dao.DB().Model(dao.Table()).Safe().Ctx(ctx)

	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		Model: model,
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
		}
	}

	model = daoctl.RegisterDaoHook(model)

	return model
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
