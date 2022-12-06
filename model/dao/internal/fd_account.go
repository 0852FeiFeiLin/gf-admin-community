// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FdAccountDao is the data access object for table fd_account.
type FdAccountDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns FdAccountColumns // columns contains all the column names of Table for convenient usage.
}

// FdAccountColumns defines and stores column names for table fd_account.
type FdAccountColumns struct {
	Id                 string // ID
	Name               string // 账户名称
	UnionLicenseId     string // 关联资质ID，大于0时必须保值与 union_user_id 关联得上
	UnionUserId        string // 关联用户ID
	CurrencyCode       string // 货币代码
	IsEnabled          string // 是否启用：1启用，0禁用
	LimitState         string // 限制状态：0不限制，1限制支出、2限制收入
	PrecisionOfBalance string // 货币单位精度：1:元，10:角，100:分，1000:厘，10000:毫，……
	Balance            string // 当前余额，必须要与账单最后一笔交易余额对应得上
	CreatedAt          string //
	UpdatedAt          string //
	DeletedAt          string //
}

// fdAccountColumns holds the columns for table fd_account.
var fdAccountColumns = FdAccountColumns{
	Id:                 "id",
	Name:               "name",
	UnionLicenseId:     "union_license_id",
	UnionUserId:        "union_user_id",
	CurrencyCode:       "currency_code",
	IsEnabled:          "is_enabled",
	LimitState:         "limit_state",
	PrecisionOfBalance: "precision_of_balance",
	Balance:            "balance",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
}

// NewFdAccountDao creates and returns a new DAO object for table data access.
func NewFdAccountDao() *FdAccountDao {
	return &FdAccountDao{
		group:   "default",
		table:   "fd_account",
		columns: fdAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FdAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FdAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FdAccountDao) Columns() FdAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FdAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FdAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FdAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
