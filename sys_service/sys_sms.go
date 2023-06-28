// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type (
	ISysSms interface {
		Verify(ctx context.Context, mobile string, captcha string, typeIdentifier ...sys_enum.SmsCaptchaType) (bool, error)
	}
)

var (
	localSysSms ISysSms
)

func SysSms() ISysSms {
	if localSysSms == nil {
		panic("implement not found for interface ISysSms, forgot register?")
	}
	return localSysSms
}

func RegisterSysSms(i ISysSms) {
	localSysSms = i
}
