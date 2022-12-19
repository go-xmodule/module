/**
 * Created by GoLand
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 16:05
 * @desc   redis.go
 */

package system

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/module/utils"
	"github.com/go-utils-module/module/utils/dirver"
	"github.com/go-utils-module/module/utils/handler"
)

// InitializeRedisPool 初始化redis连接池
func InitializeRedisPool(config config.Redis) *redis.Client {
	c, err := dirver.InitializeRedis(config.Host, config.Port, config.Password, config.Db)
	if err != nil {
		utils.Logger.Fatalln("初始化系统-连接Redis数据库异常。", err)
	}
	handler.RedisHandler = handler.NewRedis(c)
	return c
}
