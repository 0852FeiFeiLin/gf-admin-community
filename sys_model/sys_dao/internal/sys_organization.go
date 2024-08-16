// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrganizationDao is the data access object for table sys_organization.
type SysOrganizationDao struct {
	dao_interface.IDao
	table       string                 // table is the underlying table name of the DAO.
	group       string                 // group is the database configuration group name of current DAO.
	columns     SysOrganizationColumns // columns contains all the column names of Table for convenient usage.
	daoConfig   *dao_interface.DaoConfig
	ignoreCache bool
	exWhereArr  []string
}

// SysOrganizationColumns defines and stores column names for table sys_organization.
type SysOrganizationColumns struct {
	Id          string //
	Name        string // 名称
	ParentId    string // 父级ID
	CascadeDeep string // 级联深度
	Description string // 描述
}

// sysOrganizationColumns holds the columns for table sys_organization.
var sysOrganizationColumns = SysOrganizationColumns{
	Id:          "id",
	Name:        "name",
	ParentId:    "parent_id",
	CascadeDeep: "cascade_deep",
	Description: "description",
}

// NewSysOrganizationDao creates and returns a new DAO object for table data access.
func NewSysOrganizationDao(proxy ...dao_interface.IDao) *SysOrganizationDao {
	var dao *SysOrganizationDao
	if len(proxy) > 0 {
		dao = &SysOrganizationDao{
			group:       proxy[0].Group(),
			table:       proxy[0].Table(),
			columns:     sysOrganizationColumns,
			daoConfig:   proxy[0].DaoConfig(context.Background()),
			IDao:        proxy[0].DaoConfig(context.Background()).Dao,
			ignoreCache: proxy[0].DaoConfig(context.Background()).IsIgnoreCache(),
			exWhereArr:  proxy[0].DaoConfig(context.Background()).Dao.GetExtWhereKeys(),
		}

		return dao
	}

	return &SysOrganizationDao{
		group:   "default",
		table:   "sys_organization",
		columns: sysOrganizationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysOrganizationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysOrganizationDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysOrganizationDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysOrganizationDao) Columns() SysOrganizationColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysOrganizationDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SysOrganizationDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) *dao_interface.DaoConfig {
	//if dao.daoConfig != nil && len(dao.exWhereArr) == 0 {
	//	return dao.daoConfig
	//}

	var daoConfig = daoctl.NewDaoConfig(ctx, dao, cacheOption...)
	dao.daoConfig = &daoConfig

	if len(dao.exWhereArr) > 0 {
		daoConfig.IgnoreExtModel(dao.exWhereArr...)
		dao.exWhereArr = []string{}

	}

	if dao.ignoreCache {
		daoConfig.IgnoreCache()
	}

	return dao.daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysOrganizationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

func (dao *SysOrganizationDao) GetExtWhereKeys() []string {
	return dao.exWhereArr
}

func (dao *SysOrganizationDao) IsIgnoreCache() bool {
	return dao.ignoreCache
}

func (dao *SysOrganizationDao) IgnoreCache() dao_interface.IDao {
	dao.ignoreCache = true
	return dao
}
func (dao *SysOrganizationDao) IgnoreExtModel(whereKey ...string) dao_interface.IDao {
	dao.exWhereArr = append(dao.exWhereArr, whereKey...)
	return dao
}
