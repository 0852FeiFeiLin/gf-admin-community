// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"
)

type (
	ISysUser interface {
		InstallHook(state kyEnum.UserEventState, hookFunc model.UserHookFunc) int64
		UnInstallHook(savedHookId int64)
		CleanAllHook()
		QueryUserList(ctx context.Context, info *model.SearchFilter, isExport bool) (response *model.SysUserRes, err error)
		SetUserRoleIds(ctx context.Context, roleIds []int64, userId int64) (bool, error)
		CreateUser(ctx context.Context, info model.UserInnerRegister, userState kyEnum.UserState, userType kyEnum.UserType, customId ...int64) (*model.SysUserRegisterRes, error)
		GetSysUserByUsername(ctx context.Context, username string) (*entity.SysUser, error)
		HasSysUserByUsername(ctx context.Context, username string) bool
		GetSysUserById(ctx context.Context, userId int64) (*entity.SysUser, error)
		SetUserPermissionIds(ctx context.Context, userId int64, permissionIds []int64) (bool, error)
		GetUserPermissionIds(ctx context.Context, userId int64) ([]int64, error)
		SetUsername(ctx context.Context, newUsername string) (bool, error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
