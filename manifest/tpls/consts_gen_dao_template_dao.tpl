// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"{TplImportPrefix}/internal"
)

// internal{TplTableNameCamelCase}Dao is internal type for wrapping internal DAO implements.
type internal{TplTableNameCamelCase}Dao = *internal.{TplTableNameCamelCase}Dao

// {TplTableNameCamelLowerCase}Dao is the data access object for table {TplTableName}.
// You can define custom methods on it to extend its functionality as you wish.
type {TplTableNameCamelLowerCase}Dao struct {
	internal{TplTableNameCamelCase}Dao
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} = {TplTableNameCamelLowerCase}Dao{
		internal.New{TplTableNameCamelCase}Dao(),
	}
)

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *{TplTableNameCamelLowerCase}Dao) Ctx(ctx context.Context,cacheOption ...gdb.CacheOption) *gdb.Model {
	conf := gdb.CacheOption{
		Duration: time.Hour * 24,
		Force:    false,
	}

	if len(cacheOption) == 0  {
		for _, cacheConf := range sys_consts.Global.OrmCacheConf {
			if cacheConf.TableName == dao.Table() {
				conf.Duration = time.Second * (time.Duration)(cacheConf.ExpireSeconds)
				conf.Force = cacheConf.Force
			}
		}
	}else {
		conf = cacheOption[0]
	}

	return dao.DB().Model(dao.Table()).Safe().Ctx(ctx).Cache(conf)
}
