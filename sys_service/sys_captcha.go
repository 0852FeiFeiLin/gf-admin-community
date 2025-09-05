// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICaptcha interface {
		// MakeCaptcha 创建验证码，直接输出验证码图片内容到HTTP Response.
		MakeCaptcha(ctx context.Context) error
		// VerifyAndClear 校验验证码，并清空缓存的验证码信息
		VerifyAndClear(_ *ghttp.Request, value string) bool
	}
)

var (
	localCaptcha ICaptcha
)

func Captcha() ICaptcha {
	if localCaptcha == nil {
		// 记录错误但不panic，让调用方处理
		g.Log().Error(context.Background(), "ICaptcha服务未注册，请检查服务初始化顺序")
		return nil
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
