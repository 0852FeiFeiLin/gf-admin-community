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
	IMiddleware interface {
		// Auth 通讯鉴权
		Auth(r *ghttp.Request)
		// CTX 自定义上下文对象
		CTX(r *ghttp.Request)
		// CORS 允许接口跨域请求
		CORS(r *ghttp.Request)
		// ResponseHandler 响应函数
		ResponseHandler(r *ghttp.Request)
		
		// 增强安全中间件
		// EnhancedAuth 增强安全鉴权中间件
		EnhancedAuth(r *ghttp.Request)
		// SecurityMonitor 安全监控中间件
		SecurityMonitor(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		// 记录错误但不panic，让调用方处理
		g.Log().Error(context.Background(), "IMiddleware服务未注册，请检查服务初始化顺序")
		return nil
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
