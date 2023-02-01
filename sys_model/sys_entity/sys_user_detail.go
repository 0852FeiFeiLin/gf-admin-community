// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserDetail is the golang structure for table sys_user_detail.
type SysUserDetail struct {
	Id            int64       `json:"id"            description:"ID，保持与USERID一致"`
	Realname      string      `json:"realname"      description:"姓名"`
	UnionMainName string      `json:"unionMainName" description:"关联主体名称"`
	LastLoginIp   string      `json:"lastLoginIp"   description:"最后登录IP"`
	LastLoginArea string      `json:"lastLoginArea" description:"最后登录地区"`
	LastLoginAt   *gtime.Time `json:"lastLoginAt"   description:"最后登录时间"`
}
