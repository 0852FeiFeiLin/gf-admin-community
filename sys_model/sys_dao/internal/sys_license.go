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

// SysLicenseDao is the data access object for table sys_license.
type SysLicenseDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SysLicenseColumns // columns contains all the column names of Table for convenient usage.
}

// SysLicenseColumns defines and stores column names for table sys_license.
type SysLicenseColumns struct {
	Id                        string //
	IdcardFrontPath           string // 身份证头像面照片
	IdcardBackPath            string // 身份证国徽面照片
	IdcardNo                  string // 身份证号
	IdcardExpiredDate         string // 身份证有效期
	IdcardAddress             string // 身份证户籍地址
	PersonContactName         string // 负责人，必须是自然人
	PersonContactMobile       string // 负责人，联系电话
	BusinessLicenseName       string // 公司全称
	BusinessLicenseAddress    string // 公司地址
	BusinessLicensePath       string // 营业执照图片地址
	BusinessLicenseScope      string // 经营范围
	BusinessLicenseRegCapital string // 注册资本
	BusinessLicenseTermTime   string // 营业期限
	BusinessLicenseOrgCode    string // 组织机构代码
	BusinessLicenseCreditCode string // 统一社会信用代码
	BusinessLicenseLegal      string // 法人
	BusinessLicenseLegalPath  string // 法人证照，如果法人不是自然人，则该项必填
	LatestAuditLogId          string // 最新的审核记录ID
	State                     string //
	AuthType                  string //
	Remark                    string //
	UpdatedAt                 string //
	CreatedAt                 string //
	DeletedAt                 string //
}

// sysLicenseColumns holds the columns for table sys_license.
var sysLicenseColumns = SysLicenseColumns{
	Id:                        "id",
	IdcardFrontPath:           "idcard_front_path",
	IdcardBackPath:            "idcard_back_path",
	IdcardNo:                  "idcard_no",
	IdcardExpiredDate:         "idcard_expired_date",
	IdcardAddress:             "idcard_address",
	PersonContactName:         "person_contact_name",
	PersonContactMobile:       "person_contact_mobile",
	BusinessLicenseName:       "business_license_name",
	BusinessLicenseAddress:    "business_license_address",
	BusinessLicensePath:       "business_license_path",
	BusinessLicenseScope:      "business_license_scope",
	BusinessLicenseRegCapital: "business_license_reg_capital",
	BusinessLicenseTermTime:   "business_license_term_time",
	BusinessLicenseOrgCode:    "business_license_org_code",
	BusinessLicenseCreditCode: "business_license_credit_code",
	BusinessLicenseLegal:      "business_license_legal",
	BusinessLicenseLegalPath:  "business_license_legal_path",
	LatestAuditLogId:          "latest_audit_log_id",
	State:                     "state",
	AuthType:                  "auth_type",
	Remark:                    "remark",
	UpdatedAt:                 "updated_at",
	CreatedAt:                 "created_at",
	DeletedAt:                 "deleted_at",
}

// NewSysLicenseDao creates and returns a new DAO object for table data access.
func NewSysLicenseDao(proxy ...dao_interface.IDao) *SysLicenseDao {
	var dao *SysLicenseDao
	if proxy != nil {
		dao = &SysLicenseDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: sysLicenseColumns,
		}
		return dao
	}

	return &SysLicenseDao{
		group:   "default",
		table:   "sys_license",
		columns: sysLicenseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysLicenseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysLicenseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysLicenseDao) Columns() SysLicenseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysLicenseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysLicenseDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
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
func (dao *SysLicenseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
