// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysFileDao is the data access object for table sys_file.
type SysFileDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysFileColumns // columns contains all the column names of Table for convenient usage.
}

// SysFileColumns defines and stores column names for table sys_file.
type SysFileColumns struct {
	Id        string // 自增ID
	Name      string // 文件名称
	Src       string // 存储路径
	Url       string // URL地址
	Ext       string // 扩展名
	Size      string // 文件大小
	Category  string // 文件分类
	UserId    string // 用户ID
	LicenseId string // 主体ID
	CreatedAt string //
	UpdatedAt string //
}

// sysFileColumns holds the columns for table sys_file.
var sysFileColumns = SysFileColumns{
	Id:        "id",
	Name:      "name",
	Src:       "src",
	Url:       "url",
	Ext:       "ext",
	Size:      "size",
	Category:  "category",
	UserId:    "user_id",
	LicenseId: "license_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysFileDao creates and returns a new DAO object for table data access.
func NewSysFileDao() *SysFileDao {
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

// Columns returns all column names of current dao.
func (dao *SysFileDao) Columns() SysFileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
