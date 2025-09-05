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
	ISdkAliyun interface {
		// GetAliyunSdkToken 通过SDK获取Token （SDK获取方式）
		GetAliyunSdkToken(ctx context.Context, tokenInfo sys_model.AliyunSdkConfToken, err error)
		// GetAliyunSdkConfList 获取阿里云SDK应用配置列表
		GetAliyunSdkConfList(ctx context.Context) ([]*sys_model.AliyunSdkConf, error)
		// GetAliyunSdkConf 根据identifier标识获取SDK配置信息
		GetAliyunSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.AliyunSdkConf, err error)
		// SaveAliyunSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
		SaveAliyunSdkConf(ctx context.Context, info *sys_model.AliyunSdkConf, isCreate bool) (*sys_model.AliyunSdkConf, error)
		// DeleteAliyunSdkConf 删除阿里云SDK应用配置信息
		DeleteAliyunSdkConf(ctx context.Context, identifier string) (bool, error)
		// GetWsCustomizedChGeneral 中文分词
		GetWsCustomizedChGeneral(ctx context.Context, text string) (sys_model.AliyunNlpDataRes, error)
	}
)

var (
	localSdkAliyun ISdkAliyun
)

func SdkAliyun() ISdkAliyun {
	if localSdkAliyun == nil {
		// 记录错误但不panic，让调用方处理
		g.Log().Error(context.Background(), "ISdkAliyun服务未注册，请检查服务初始化顺序")
		return nil
	}
	return localSdkAliyun
}

func RegisterSdkAliyun(i ISdkAliyun) {
	localSdkAliyun = i
}
