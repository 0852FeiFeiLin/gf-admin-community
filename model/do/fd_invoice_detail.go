// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FdInvoiceDetail is the golang structure of table fd_invoice_detail for DAO operations like Where/Data.
type FdInvoiceDetail struct {
	g.Meta        `orm:"table:fd_invoice_detail, do:true"`
	Id            interface{} // ID
	TaxNumber     interface{} // 纳税识别号
	TaxName       interface{} // 纳税人名称
	BillIds       interface{} // 账单ID组
	Amount        interface{} // 开票金额，单位精度：分
	Rate          interface{} // 税率，如3% 则填入3
	RateMount     interface{} // 税额，单位精度：分
	Remark        interface{} // 发布内容描述
	Type          interface{} // 发票类型：1电子发票，2纸质发票
	State         interface{} // 状态：1待审核、2待开票、4开票失败、8已开票、16已撤销
	AuditUserIds  interface{} // 审核者UserID，多个用逗号隔开
	MakeType      interface{} // 出票类型：1普通发票、2增值税专用发票、3专业发票
	MakeUserId    interface{} // 出票人UserID，如果是系统出票则默认-1
	MakeAt        *gtime.Time // 出票时间
	CourierName   interface{} // 快递名称
	CourierNumber interface{} // 快递编号
	FdInvoiceId   interface{} // 发票抬头ID
	AuditUserId   interface{} // 审核者UserID
	AuditReplyMsg interface{} // 审核回复，仅审核不通过时才有值
	AuditAt       *gtime.Time // 审核时间
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	DeletedAt     *gtime.Time //
	UserId        interface{} // 申请者用户ID
	UnionMainId   interface{} // 主体ID：运营商ID、服务商ID、商户ID、消费者ID
}
