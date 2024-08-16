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

// SysFileDao is the data access object for table sys_file.
type SysFileDao struct {
	dao_interface.IDao
	table       string         // table is the underlying table name of the DAO.
	group       string         // group is the database configuration group name of current DAO.
	columns     SysFileColumns // columns contains all the column names of Table for convenient usage.
	daoConfig   *dao_interface.DaoConfig
	ignoreCache bool
	exWhereArr  []string
}

// SysFileColumns defines and stores column names for table sys_file.
type SysFileColumns struct {
	Id          string // 自增ID
	Name        string // 文件名称
	Src         string // 存储路径
	Url         string // URL地址
	Ext         string // 扩展名
	Size        string // 文件大小
	Category    string // 文件分类
	UserId      string // 用户ID
	UnionMainId string // 关联主体ID
	CreatedAt   string //
	UpdatedAt   string //
	LocalPath   string // 本地路径
}

// sysFileColumns holds the columns for table sys_file.
var sysFileColumns = SysFileColumns{
	Id:          "id",
	Name:        "name",
	Src:         "src",
	Url:         "url",
	Ext:         "ext",
	Size:        "size",
	Category:    "category",
	UserId:      "user_id",
	UnionMainId: "union_main_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	LocalPath:   "local_path",
}

// NewSysFileDao creates and returns a new DAO object for table data access.
func NewSysFileDao(proxy ...dao_interface.IDao) *SysFileDao {
	var dao *SysFileDao
	if len(proxy) > 0 {
		dao = &SysFileDao{
			group:       proxy[0].Group(),
			table:       proxy[0].Table(),
			columns:     sysFileColumns,
			daoConfig:   proxy[0].DaoConfig(context.Background()),
			IDao:        proxy[0].DaoConfig(context.Background()).Dao,
			ignoreCache: proxy[0].DaoConfig(context.Background()).IsIgnoreCache(),
			exWhereArr:  proxy[0].DaoConfig(context.Background()).Dao.GetExtWhereKeys(),
		}

		return dao
	}

	return &SysFileDao{
		group:   "default",
		table:   "sys_file",
		columns: sysFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysFileDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysFileDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysFileDao) Columns() SysFileColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysFileDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SysFileDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) *dao_interface.DaoConfig {
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
func (dao *SysFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

func (dao *SysFileDao) GetExtWhereKeys() []string {
	return dao.exWhereArr
}

func (dao *SysFileDao) IsIgnoreCache() bool {
	return dao.ignoreCache
}

func (dao *SysFileDao) IgnoreCache() dao_interface.IDao {
	dao.ignoreCache = true
	return dao
}
func (dao *SysFileDao) IgnoreExtModel(whereKey ...string) dao_interface.IDao {
	dao.exWhereArr = append(dao.exWhereArr, whereKey...)
	return dao
}
