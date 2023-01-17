package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
)

// SysUser 鉴权
var SysUser = cSysUser{}

type cSysUser struct{}

// QueryUserList 获取用户|列表
func (c *cSysUser) QueryUserList(ctx context.Context, req *sys_api.QueryUserListReq) (*sys_model.SysUserListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*sys_model.SysUserListRes, error) {
			sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
			return sys_service.SysUser().QueryUserList(
				ctx,
				&req.SearchParams,
				sessionUser.UnionMainId,
				false,
			)
		},
		sys_enum.User.PermissionType.List,
	)
}

// SetUserPermissionIds 设置用户权限
func (c *cSysUser) SetUserPermissionIds(ctx context.Context, req *sys_api.SetUserPermissionIdsReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := sys_service.SysUser().SetUserPermissionIds(
				ctx,
				req.Id,
				req.PermissionIds,
			)
			return ret == true, err
		},
		sys_enum.User.PermissionType.SetPermission,
	)
}

// GetUserPermissionIds 获取用户权限Ids
func (c *cSysUser) GetUserPermissionIds(ctx context.Context, req *sys_api.GetUserPermissionIdsReq) (api_v1.Int64ArrRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.Int64ArrRes, error) {
			return sys_service.SysPermission().GetPermissionsByResource(
				ctx,
				req.Id,
			)
		},
		sys_enum.User.PermissionType.SetPermission,
	)
}

// ResetUserPassword 重置用户密码
func (c *cSysUser) ResetUserPassword(ctx context.Context, req *sys_api.ResetUserPasswordReq) (res api_v1.BoolRes, err error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			// 获取当前登录用户
			user := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := sys_service.SysUser().ResetUserPassword(
				ctx,
				req.UserId,
				req.Password,
				req.ConfirmPassword,
				user.SysUser,
			)
			return ret == true, err
		},
		sys_enum.User.PermissionType.ResetPassword,
	)
}
