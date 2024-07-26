package sdk_tencent

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

// 腾讯云服务平台

type sSdkTencent struct {
	TencentSdkConfTokenList []*sys_model.TencentSdkConfToken
	sysConfigName           string
	conf                    gdb.CacheOption
}

// New SdkBaidu 系统配置逻辑实现
func New() sys_service.ISdkTencent {
	return &sSdkTencent{
		TencentSdkConfTokenList: make([]*sys_model.TencentSdkConfToken, 0),
		sysConfigName:           "tencent_sdk_conf",
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

func init() {
	sys_service.RegisterSdkTencent(New())
}

// fetchTencentSdkConfToken 根据 identifier 获取腾讯云API Token  （API获取方式）
func (s *sSdkTencent) fetchTencentSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConfToken, err error) {

	info, err := s.GetTencentSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}
	client := g.Client()

	// URL 请求的服务器URL
	var host = "https://rkp.tencentcloudapi.com"

	// 请求头
	header := make(map[string]string)

	header["X-TC-Action"] = "GetToken"
	header["Content-type"] = "application/json"
	header["X-TC-Region"] = ""
	header["X-TC-Timestamp"] = gtime.Now().TimestampStr()
	header["X-TC-Version"] = info.Version
	// header["Authorization"] = ""
	header["X-TC-Language"] = "zh-CN"

	client.Header(header)

	// 请求数据，
	param := g.Map{
		// 业务ID
		"BusinessId": gconv.Int64(info.AppID),
		// 业务子场景
		"Scene": 0,
		// 业务侧账号体系下的用户ID (不是必填)
		"BusinessUserId": info.AESKey,
		// 用户侧的IP (不是必填)
		"AppClientIp": info.AppID,
		// 过期时间 (不是必填)
		"ExpireTime": info.APIKey,
	}

	response, err := client.Post(ctx, host, param)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取腾讯云API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 接受返回数据，json解析
	newTokenInfo := sys_model.TencentAccessToken{}

	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取腾讯云API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	var result *sys_model.TencentSdkConfToken = nil
	newItems := garray.New()
	for _, item := range s.TencentSdkConfTokenList {
		if item.Identifier == identifier {
			result = &sys_model.TencentSdkConfToken{
				TencentSdkConf:     *info,
				TencentAccessToken: newTokenInfo,
			}
			newItems.Append(*result)
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &sys_model.TencentSdkConfToken{
			TencentSdkConf:     *info,
			TencentAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	// 返回我们需要的token信息
	return result, nil
}

// GetTencentSdkConfList 获取腾讯云SDK应用配置列表
func (s *sSdkTencent) GetTencentSdkConfList(ctx context.Context) ([]*sys_model.TencentSdkConf, error) {
	items := make([]*sys_model.TencentSdkConf, 0)
	config, err := sys_service.SysConfig().GetByName(ctx, s.sysConfigName)
	if err != nil {
		return items, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("腾讯云SDK配置信息获取失败"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if config.Value == "" {
		return items, nil
	}

	_ = gjson.DecodeTo(config.Value, &items)

	return items, nil
}

// GetTencentSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkTencent) GetTencentSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConf, err error) {
	items, err := s.GetTencentSdkConfList(ctx)
	if err != nil {
		return nil, err
	}

	// 循环所有配置，筛选出符合条件的配置
	for _, conf := range items {
		if conf.Identifier == identifier {
			return conf, nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询腾讯云SDK应用配置信息失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveTencentSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkTencent) SaveTencentSdkConf(ctx context.Context, info *sys_model.TencentSdkConf, isCreate bool) (*sys_model.TencentSdkConf, error) {
	oldItems, _ := s.GetTencentSdkConfList(ctx)

	isHas := false
	newItems := make([]*sys_model.TencentSdkConf, 0)
	for _, conf := range oldItems {
		if conf.Identifier == info.Identifier { // 如果标识符相等，说明已经存在， 将最新的追加到新的容器中
			isHas = true
			newItems = append(newItems, info)
			continue
		}

		newItems = append(newItems, conf) // 将旧的Item追加到新的容器中
	}

	if !isHas { // 不存在
		if isCreate { // 创建 --- 追加info （原有的 + 最新的Info）
			newItems = append(newItems, info)
		} else { // 更新 --- 不存在此配置，那么就提示错误
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("腾讯云SDK配置信息保存失败，标识符错误"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	// 序列化后进行保存至数据库
	jsonString := gjson.MustEncodeString(newItems)
	_, err := sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "腾讯云SDK配置信息保存失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步腾讯云SDK应用配置缓存列表
	s.syncTencentSdkConfList(ctx)

	return info, nil
}

// syncTencentSdkConfList 同步腾讯云SDK应用配置信息列表缓存  （代码中要是用到了s.TencentSdkConfList缓存变量的话，一定需要在CUD操作后调用此方法更新缓存变量）
func (s *sSdkTencent) syncTencentSdkConfList(ctx context.Context) error {
	items, err := s.GetTencentSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]*sys_model.TencentSdkConfToken, 0)
	for _, conf := range items {
		for _, tokenInfo := range s.TencentSdkConfTokenList { // tokenList
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.TencentSdkConfTokenList = newTokenItems

	return nil
}

// DeleteTencentSdkConf 删除腾讯云SDK应用配置信息
func (s *sSdkTencent) DeleteTencentSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetTencentSdkConfList(ctx)

	isHas := false
	newItems := garray.New(false)
	for _, conf := range items {
		if conf.Identifier == identifier {
			isHas = true
			continue
		}
		newItems.Append(conf)
	}

	if !isHas {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "要删除的腾讯云SDK配置信息不存在", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)
	_, err = sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "腾讯云SDK配置信息删除失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步Token列表
	s.syncTencentSdkConfList(ctx)

	return true, nil
}
