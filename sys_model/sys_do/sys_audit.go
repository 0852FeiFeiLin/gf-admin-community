// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAudit is the golang structure of table sys_audit for DAO operations like Where/Data.
type SysAudit struct {
	g.Meta       `orm:"table:sys_audit, do:true"`
	Id           interface{} //
	State        interface{} // 审核状态：-1不通过，0待审核，1通过
	Reply        interface{} // 不通过时回复的审核不通过原因
	UnionMainId  interface{} // 关联主体ID
	Category     interface{} // 业务类别
	AuditData    interface{} // 待审核的业务数据包
	ExpireAt     *gtime.Time // 服务时限
	AuditReplyAt *gtime.Time // 审核回复时间
	HistoryItems interface{} // 历史申请记录
	CreatedAt    *gtime.Time //
	AuditUserId  interface{} // 审核操作者id
}
