// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICasbin interface {
		// InstallHook 安装Hook
		InstallHook(userType sys_enum.UserType, hookFunc sys_hook.CasbinHookFunc) int64
		// UnInstallHook 卸载Hook
		UnInstallHook(savedHookId int64)
		// CleanAllHook 清除所有Hook
		CleanAllHook()
		Check() error
		Enforcer() *casbin.Enforcer
		// AddRoleForUserInDomain 添加用户角色关联关系
		AddRoleForUserInDomain(userName string, roleName string, domain string) (bool, error)
		// DeleteRoleForUserInDomain 删除用户角色关联关系
		DeleteRoleForUserInDomain(userName string, roleName string, domain string) (bool, error)
		// DeleteRolesForUser 清空用户角色关联关系
		DeleteRolesForUser(userName string, domain string) (bool, error)
		// AddPermissionForUser 添加角色与资源关系
		AddPermissionForUser(roleName string, path string, method string) (bool, error)
		// AddPermissionsForUser 添加角色与资源关系
		AddPermissionsForUser(roleName string, path []string) (bool, error)
		// DeletePermissionForUser 删除角色与资源关系
		DeletePermissionForUser(roleName string, path string, method string) (bool, error)
		// DeletePermissionsForUser 清空角色与资源关系
		DeletePermissionsForUser(roleName string) (bool, error)
		// EnforceCheck 校验  确认访问权限
		EnforceCheck(userName interface{}, path interface{}, role interface{}, method interface{}) (bool, error)
	}
)

var (
	localCasbin ICasbin
)

func Casbin() ICasbin {
	if localCasbin == nil {
		// 记录错误但不panic，让调用方处理
		g.Log().Error(context.Background(), "ICasbin服务未注册，请检查服务初始化顺序")
		return nil
	}
	return localCasbin
}

func RegisterCasbin(i ICasbin) {
	localCasbin = i
}
