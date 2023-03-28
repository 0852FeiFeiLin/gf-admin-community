// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISessionError interface {
		Append(ctx context.Context, error error) error
		HasError(ctx context.Context, err error) (response bool)
		Iterator(ctx context.Context, f func(k int, err error) bool)
	}
	ISysSession interface {
		Init(sessionContext *sys_model.SessionContext, r *ghttp.Request, ctx ...*context.Context)
		NewSessionCtx(ctx context.Context) context.Context
		HasCustom(ctx context.Context) bool
		Get(ctx context.Context) *sys_model.SessionContext
		SetUserById(ctx *context.Context, userId int64) *sys_model.SessionContext
		SetUser(ctx context.Context, claimsUser *sys_model.JwtCustomClaims)
	}
)

var (
	localSessionError ISessionError
	localSysSession   ISysSession
)

func SessionError() ISessionError {
	if localSessionError == nil {
		panic("implement not found for interface ISessionError, forgot register?")
	}
	return localSessionError
}

func RegisterSessionError(i ISessionError) {
	localSessionError = i
}

func SysSession() ISysSession {
	if localSysSession == nil {
		panic("implement not found for interface ISysSession, forgot register?")
	}
	return localSysSession
}

func RegisterSysSession(i ISysSession) {
	localSysSession = i
}
