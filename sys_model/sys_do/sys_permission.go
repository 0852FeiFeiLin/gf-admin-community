// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure of table sys_permission for DAO operations like Where/Data.
type SysPermission struct {
	g.Meta      `orm:"table:sys_permission, do:true"`
	Id          interface{} // ID
	ParentId    interface{} // 父级ID
	Name        interface{} // 名称
	Description interface{} // 描述
	Identifier  interface{} // 标识符
	Type        interface{} // 类型：1api，2menu
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	IsShow      interface{} // 是否显示：0不显示 1显示
	Sort        interface{} // 排序字段
}
