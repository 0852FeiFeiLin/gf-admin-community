package sys_permission

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sSysPermission struct {
	CacheDuration time.Duration
	CachePrefix   string
}

func init() {
	sys_service.RegisterSysPermission(New())
}

// New sSysPermission 权限控制逻辑实现
func New() *sSysPermission {
	return &sSysPermission{
		CacheDuration: time.Hour,
		CachePrefix:   sys_dao.SysPermission.Table() + "_",
	}
}

// GetPermissionById 根据权限ID获取权限信息
func (s *sSysPermission) GetPermissionById(ctx context.Context, permissionId int64) (*sys_entity.SysPermission, error) {
	result := sys_entity.SysPermission{}

	err := sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{Id: permissionId}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", sys_dao.SysPermission.Table())
	}

	return &result, nil
}

// GetPermissionByName 根据权限Name获取权限信息
func (s *sSysPermission) GetPermissionByName(ctx context.Context, permissionName string) (*sys_entity.SysPermission, error) {
	result := sys_entity.SysPermission{}

	err := sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{Name: permissionName}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", sys_dao.SysPermission.Table())
	}

	return &result, nil
}

// QueryPermissionList 查询权限
func (s *sSysPermission) QueryPermissionList(ctx context.Context, info sys_model.SearchParams) (*sys_model.SysPermissionInfoListRes, error) {
	result, err := daoctl.Query[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx), &info, false)
	return (*sys_model.SysPermissionInfoListRes)(result), err
}

// GetPermissionList 根据ID获取下级权限信息，返回列表
func (s *sSysPermission) GetPermissionList(ctx context.Context, parentId int64, IsRecursive bool) (*[]sys_entity.SysPermission, error) {
	result := make([]sys_entity.SysPermission, 0)
	err := sys_dao.SysPermission.Ctx(ctx).
		// 数据查询结果缓存起来
		Cache(gdb.CacheOption{
			Duration: s.CacheDuration,
			Name:     s.CachePrefix + gconv.String(parentId),
			Force:    true,
		}).
		Where(sys_do.SysPermission{
			ParentId: parentId,
			IsShow:   1,
		}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	}

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(result) > 0 {
		for _, sysPermissionItem := range result {
			var children *[]sys_entity.SysPermission
			children, err = s.GetPermissionList(ctx, sysPermissionItem.Id, IsRecursive)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
			}

			if children == nil || len(*children) <= 0 {
				continue
			}

			for _, sysOrganization := range *children {
				result = append(result, sysOrganization)
			}
		}
	}

	return &result, nil
}

// GetPermissionTree 根据ID获取下级权限信息，返回列表树
func (s *sSysPermission) GetPermissionTree(ctx context.Context, parentId int64) (*[]sys_model.SysPermissionTree, error) {
	result, err := s.GetPermissionList(ctx, parentId, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	}

	response := make([]sys_model.SysPermissionTree, 0)

	// 有数据，则递归加载
	if len(*result) > 0 {
		for _, sysPermissionItem := range *result {
			item := sys_model.SysPermissionTree{}
			gconv.Struct(sysPermissionItem, &item)

			item.Children, err = s.GetPermissionTree(ctx, sysPermissionItem.Id)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
			}

			response = append(response, item)
		}
	}
	return &response, nil
}

func (s *sSysPermission) CreatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	return s.SavePermission(ctx, info)
}

func (s *sSysPermission) UpdatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", sys_dao.SysPermission.Table())
	}
	return s.SavePermission(ctx, info)
}

// ImportPermission 导入权限
func (s *sSysPermission) ImportPermission(ctx context.Context, infoArr []*sys_model.SysPermission) error {
	if len(infoArr) <= 0 {
		return nil
	}

	idArr := make([]int64, 0)
	idMap := gmap.New()

	// 提取所有id
	for _, permission := range infoArr {
		idArr = append(idArr, permission.Id)
		idMap.Set(permission.Id, permission)
	}

	data := make([]sys_entity.SysPermission, 0)
	// 加载已导入的权限
	err := sys_dao.SysPermission.Ctx(ctx).WhereIn(sys_dao.SysPermission.Columns().Id, idArr).Scan(&data)
	if err != nil {
		return err
	}

	all := make([]sys_entity.SysPermission, 0)
	// 加载所有权限
	err = sys_dao.SysPermission.Ctx(ctx).Scan(&all)
	if err != nil {
		return err
	}

	waitSaveArr := make([]*sys_model.SysPermission, 0)
	for _, permission := range infoArr {
		if idMap.Contains(permission.Id) {
			continue
		}

		if permission.ParentId > 0 {
			has := false
			for _, sysPermission := range all {
				if sysPermission.Id == permission.ParentId {
					has = true
				}
			}
			for _, sysPermission := range infoArr {
				if sysPermission.Id == permission.ParentId {
					has = true
				}
			}
			if has == false {
				fmt.Printf("权限ID父级无效[%v]：%v | %s，%s\n", permission.Id, permission.Name, permission.Identifier, permission.Description)
				continue
			}
		}
		waitSaveArr = append(waitSaveArr, permission)
	}

	for _, permission := range waitSaveArr {
		_, err = sys_dao.SysPermission.Ctx(ctx).Insert(permission)
		if err != nil {
			fmt.Printf("权限导入失败[%v]：%v | %s，%s\n", permission.Id, permission.Name, permission.Identifier, permission.Description)
			fmt.Printf("失败原因：%s\n", err.Error())
		}
	}
	// 移除已缓存的数据
	daoctl.RemoveQueryCache(sys_dao.SysPermission.DB(), s.CachePrefix)
	return nil
}

// SavePermission 新增/保存权限信息
func (s *sSysPermission) SavePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	data := sys_entity.SysPermission{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级权限信息是否存在
	if data.ParentId > 0 {
		permissionInfo, err := s.GetPermissionById(ctx, data.ParentId)
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "父级权限信息不存在"), "", sys_dao.SysPermission.Table())
		}
	}

	if data.Id <= 0 {
		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := sys_dao.SysPermission.Ctx(ctx).Insert(data)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增权限信息失败", sys_dao.SysPermission.Table())
		}
	} else {
		data.UpdatedAt = gtime.Now()
		_, err := sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{Id: data.Id}).Update(sys_do.SysPermission{
			ParentId:    data.ParentId,
			Name:        data.Name,
			Description: data.Description,
			Identifier:  data.Identifier,
			IsShow:      1,
			Type:        data.Type,
		})

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息保存失败", sys_dao.SysPermission.Table())
		}
	}

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(sys_dao.SysPermission.DB(), s.CachePrefix)
	return &data, nil
}

// DeletePermission 删除权限信息
func (s *sSysPermission) DeletePermission(ctx context.Context, permissionId int64) (bool, error) {
	_, err := s.GetPermissionById(ctx, permissionId)

	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysPermission.Ctx(ctx).Delete(sys_do.SysPermission{Id: permissionId})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除权限信息失败", sys_dao.SysPermission.Table())
	}

	// 删除权限定义
	sys_dao.SysCasbin.Ctx(ctx).Delete(sys_do.SysCasbin{Ptype: "p", V2: permissionId})

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(sys_dao.SysPermission.DB(), s.CachePrefix)
	return true, nil
}

// GetPermissionTreeIdByUrl 根据请求URL去匹配权限树，返回权限
func (s *sSysPermission) GetPermissionTreeIdByUrl(ctx context.Context, path string) (*sys_entity.SysPermission, error) {
	if path == "" {
		return nil, gerror.New("传入的请求url为空")
	}

	result := sys_entity.SysPermission{}

	// 在权限树标识中匹标识后缀，|为标识符的分隔符
	path = "%|" + path

	err := sys_dao.SysPermission.Ctx(ctx).WhereLike(sys_dao.SysPermission.Columns().Identifier, path).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CheckPermission 校验权限
func (s *sSysPermission) CheckPermission(ctx context.Context, permissionId interface{}) (bool, error) { // 权限id  域 资源  方法
	user := sys_service.BizCtx().Get(ctx).ClaimsUser.SysUser
	t, err := sys_service.Casbin().Enforcer().Enforce(user.Id, sys_consts.CasbinDomain, permissionId, "allow")
	if err != nil {
		fmt.Printf("权限校验失败[%v]：%v\n", permissionId, err.Error())
	}
	if t != true {
		err = gerror.New("没有权限")
	}
	return t, err
}

// NewPermission 构造权限信息
func (s *sSysPermission) NewPermission(code int64, identifier string, name string, description string, parentId ...int64) *sys_model.SysPermission {
	var pid int64 = 0

	if len(parentId) > 0 {
		pid = parentId[0]
	}

	result := &sys_model.SysPermission{
		Id:          code,
		ParentId:    pid,
		Name:        name,
		Description: description,
		Identifier:  identifier,
		Type:        1,
	}
	return result
}

// PermissionTypeForm 通过枚举值取枚举类型
func (s *sSysPermission) PermissionTypeForm(code int64, mapItems *gmap.StrAnyMap) *sys_model.SysPermission {
	var result *sys_model.SysPermission
	mapItems.Iterator(func(k string, v interface{}) bool {
		item := v.(*sys_model.SysPermission)
		if item.Id == code {
			result = item
			return false
		}
		return true
	})

	return result
}

// SavePermissionToDb 保存限信息到数据库，如果存在会自动忽略
func (s *sSysPermission) SavePermissionToDb(mapItems *gmap.StrAnyMap) {
	if mapItems.Size() <= 0 {
		return
	}

	waitSaveArr := make([]*sys_model.SysPermission, 0)
	mapItems.Iterator(func(k string, v interface{}) bool {
		item := v.(*sys_model.SysPermission)
		waitSaveArr = append(waitSaveArr, item)
		return true
	})

	sys_service.SysPermission().ImportPermission(context.Background(), waitSaveArr)
}
