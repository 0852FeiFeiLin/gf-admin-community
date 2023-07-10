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

// SysPersonAuditDao is the data access object for table sys_person_audit.
type SysPersonAuditDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysPersonAuditColumns // columns contains all the column names of Table for convenient usage.
}

// SysPersonAuditColumns defines and stores column names for table sys_person_audit.
type SysPersonAuditColumns struct {
	Id           string //
	State        string // 审核状态：-1不通过，0待审核，1通过
	Reply        string // 不通过时回复的审核不通过原因
	UnionMainId  string // 关联主体ID
	Category     string // 业务类别
	AuditData    string // 待审核的业务数据包
	ExpireAt     string // 服务时限
	AuditReplyAt string // 审核回复时间
	HistoryItems string // 历史申请记录
	CreatedAt    string //
	AuditUserId  string // 审核操作者id
}

// sysPersonAuditColumns holds the columns for table sys_person_audit.
var sysPersonAuditColumns = SysPersonAuditColumns{
	Id:           "id",
	State:        "state",
	Reply:        "reply",
	UnionMainId:  "union_main_id",
	Category:     "category",
	AuditData:    "audit_data",
	ExpireAt:     "expire_at",
	AuditReplyAt: "audit_reply_at",
	HistoryItems: "history_Items",
	CreatedAt:    "created_at",
	AuditUserId:  "audit_user_id",
}

// NewSysPersonAuditDao creates and returns a new DAO object for table data access.
func NewSysPersonAuditDao(proxy ...dao_interface.IDao) *SysPersonAuditDao {
	var dao *SysPersonAuditDao
	if len(proxy) > 0 {
		dao = &SysPersonAuditDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: sysPersonAuditColumns,
		}
		return dao
	}

	return &SysPersonAuditDao{
		group:   "default",
		table:   "sys_person_audit",
		columns: sysPersonAuditColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPersonAuditDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPersonAuditDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPersonAuditDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysPersonAuditDao) Columns() SysPersonAuditColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPersonAuditDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SysPersonAuditDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysPersonAuditDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
