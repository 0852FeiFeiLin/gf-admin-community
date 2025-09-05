package sys_service

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

// ServiceValidator 服务注册验证器
type ServiceValidator struct{}

var Validator = &ServiceValidator{}

// ValidateServices 验证所有服务是否已正确注册
func (v *ServiceValidator) ValidateServices() error {
	var errors []string

	// 检查所有必需的服务是否已注册
	if localSysUser == nil {
		errors = append(errors, "ISysUser服务未注册")
	}
	if localSysAuth == nil {
		errors = append(errors, "ISysAuth服务未注册")
	}
	if localSysLogs == nil {
		errors = append(errors, "ISysLogs服务未注册")
	}
	if localSysPermission == nil {
		errors = append(errors, "ISysPermission服务未注册")
	}
	if localSdkBaidu == nil {
		errors = append(errors, "ISdkBaidu服务未注册")
	}
	if localSdkAliyun == nil {
		errors = append(errors, "ISdkAliyun服务未注册")
	}
	if localSdkCtyun == nil {
		errors = append(errors, "ISdkCtyun服务未注册")
	}
	if localSysMenu == nil {
		errors = append(errors, "ISysMenu服务未注册")
	}
	if localJwt == nil {
		errors = append(errors, "IJwt服务未注册")
	}
	if localSysOrganization == nil {
		errors = append(errors, "ISysOrganization服务未注册")
	}
	if localSysRole == nil {
		errors = append(errors, "ISysRole服务未注册")
	}
	if localCaptcha == nil {
		errors = append(errors, "ICaptcha服务未注册")
	}
	if localFile == nil {
		errors = append(errors, "IFile服务未注册")
	}
	if localBizCtx == nil {
		errors = append(errors, "IBizCtx服务未注册")
	}
	if localCasbin == nil {
		errors = append(errors, "ICasbin服务未注册")
	}
	if localSdkHuawei == nil {
		errors = append(errors, "ISdkHuawei服务未注册")
	}
	if localSdkTencent == nil {
		errors = append(errors, "ISdkTencent服务未注册")
	}
	if localMiddleware == nil {
		errors = append(errors, "IMiddleware服务未注册")
	}
	if localSysSms == nil {
		errors = append(errors, "ISysSms服务未注册")
	}
	if localArea == nil {
		errors = append(errors, "IArea服务未注册")
	}

	if len(errors) > 0 {
		return fmt.Errorf("服务注册验证失败: %s", strings.Join(errors, ", "))
	}
	return nil
}

// ValidateAndLog 验证服务并记录结果
func (v *ServiceValidator) ValidateAndLog(ctx context.Context) {
	if err := v.ValidateServices(); err != nil {
		g.Log().Warning(ctx, "服务验证失败", err)
	} else {
		g.Log().Info(ctx, "所有服务验证通过")
	}
}

// GetSafeService 安全获取服务，如果服务未注册则返回错误而不是panic
func (v *ServiceValidator) GetSafeService(serviceName string, service interface{}) error {
	if service == nil {
		return fmt.Errorf("服务 %s 未注册，请检查服务初始化顺序", serviceName)
	}
	return nil
}