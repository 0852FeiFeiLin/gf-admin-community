// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type (
	ISysPermission interface {
		GetPermissionById(ctx context.Context, permissionId int64) (*sys_entity.SysPermission, error)
		GetPermissionByName(ctx context.Context, permissionName string) (*sys_entity.SysPermission, error)
		QueryPermissionList(ctx context.Context, info sys_model.SearchParams) (*sys_model.SysPermissionInfoListRes, error)
		GetPermissionList(ctx context.Context, parentId int64, IsRecursive bool) (*[]sys_entity.SysPermission, error)
		GetPermissionTree(ctx context.Context, parentId int64) (*[]sys_model.SysPermissionTree, error)
		CreatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error)
		UpdatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error)
		SavePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error)
		DeletePermission(ctx context.Context, permissionId int64) (bool, error)
		GetPermissionTreeIdByUrl(ctx context.Context, path string) (int64, error)
	}
)

var (
	localSysPermission ISysPermission
)

func SysPermission() ISysPermission {
	if localSysPermission == nil {
		panic("implement not found for interface ISysPermission, forgot register?")
	}
	return localSysPermission
}

func RegisterSysPermission(i ISysPermission) {
	localSysPermission = i
}
