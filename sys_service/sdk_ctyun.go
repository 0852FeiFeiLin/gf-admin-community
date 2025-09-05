// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	ISdkCtyun interface {
		// GetCtyunSdkConfList 获取天翼云SDK应用配置列表
		GetCtyunSdkConfList(ctx context.Context) ([]*sys_model.CtyunSdkConf, error)
		// GetCtyunSdkConf 根据identifier标识获取SDK配置信息
		GetCtyunSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.CtyunSdkConf, err error)
		// SaveCtyunSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
		SaveCtyunSdkConf(ctx context.Context, info *sys_model.CtyunSdkConf, isCreate bool) (*sys_model.CtyunSdkConf, error)
		// DeleteCtyunSdkConf 删除天翼云SDK应用配置信息
		DeleteCtyunSdkConf(ctx context.Context, identifier string) (bool, error)
	}
)

var (
	localSdkCtyun ISdkCtyun
)

func SdkCtyun() ISdkCtyun {
	if localSdkCtyun == nil {
		// 记录错误但不panic，让调用方处理
		g.Log().Error(context.Background(), "ISdkCtyun服务未注册，请检查服务初始化顺序")
		return nil
	}
	return localSdkCtyun
}

func RegisterSdkCtyun(i ISdkCtyun) {
	localSdkCtyun = i
}
