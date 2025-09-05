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
	ISdkHuawei interface {
		// GetHuaweiSdkConfList 获取阿里云SDK应用配置列表
		GetHuaweiSdkConfList(ctx context.Context) ([]*sys_model.HuaweiSdkConf, error)
		// GetHuaweiSdkConf 根据identifier标识获取SDK配置信息
		GetHuaweiSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.HuaweiSdkConf, err error)
		// SaveHuaweiSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
		SaveHuaweiSdkConf(ctx context.Context, info *sys_model.HuaweiSdkConf, isCreate bool) (*sys_model.HuaweiSdkConf, error)
		// DeleteHuaweiSdkConf 删除华为SDK应用配置信息
		DeleteHuaweiSdkConf(ctx context.Context, identifier string) (bool, error)
	}
)

var (
	localSdkHuawei ISdkHuawei
)

func SdkHuawei() ISdkHuawei {
	if localSdkHuawei == nil {
		// 记录错误但不panic，让调用方处理
		g.Log().Error(context.Background(), "ISdkHuawei服务未注册，请检查服务初始化顺序")
		return nil
	}
	return localSdkHuawei
}

func RegisterSdkHuawei(i ISdkHuawei) {
	localSdkHuawei = i
}
