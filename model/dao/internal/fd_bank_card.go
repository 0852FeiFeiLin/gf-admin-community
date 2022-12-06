// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FdBankCardDao is the data access object for table fd_bank_card.
type FdBankCardDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns FdBankCardColumns // columns contains all the column names of Table for convenient usage.
}

// FdBankCardColumns defines and stores column names for table fd_bank_card.
type FdBankCardColumns struct {
	Id            string // ID
	BankName      string // 银行名称
	CardType      string // 银行卡类型：1借记卡，2储蓄卡
	CardNumber    string // 银行卡号
	ExpiredAt     string // 有效期
	HolderName    string // 银行卡开户名
	BankOfAccount string // 开户行
	State         string // 状态：0禁用，1正常
	Remark        string // 备注信息
	CreatedAt     string //
	UpdatedAt     string //
	DeletedAt     string //
}

// fdBankCardColumns holds the columns for table fd_bank_card.
var fdBankCardColumns = FdBankCardColumns{
	Id:            "id",
	BankName:      "bank_name",
	CardType:      "card_type",
	CardNumber:    "card_number",
	ExpiredAt:     "expired_at",
	HolderName:    "holder_name",
	BankOfAccount: "bank_of_account",
	State:         "state",
	Remark:        "remark",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewFdBankCardDao creates and returns a new DAO object for table data access.
func NewFdBankCardDao() *FdBankCardDao {
	return &FdBankCardDao{
		group:   "default",
		table:   "fd_bank_card",
		columns: fdBankCardColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FdBankCardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FdBankCardDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FdBankCardDao) Columns() FdBankCardColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FdBankCardDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FdBankCardDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FdBankCardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
