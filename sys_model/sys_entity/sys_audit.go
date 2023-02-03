// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAudit is the golang structure for table sys_audit.
type SysAudit struct {
	Id            int64       `json:"id"            description:""`
	State         int         `json:"state"         description:"审核状态：-1不通过，0待审核，1通过"`
	Replay        string      `json:"replay"        description:"不通过时回复的审核不通过原因"`
	UnionMainId   int64       `json:"unionMainId"   description:"关联主体ID"`
	Category      int         `json:"category"      description:"业务类别"`
	AuditData     string      `json:"auditData"     description:"待审核的业务数据包"`
	ExpireAt      *gtime.Time `json:"expireAt"      description:"服务时限"`
	AuditReplayAt *gtime.Time `json:"auditReplayAt" description:"审核回复时间"`
	HistoryItems  string      `json:"historyItems"  description:"历史申请记录"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
}
