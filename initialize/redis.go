/**
 * Created by GoLand
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 16:05
 * @desc   redis.go
 */

package system

import (
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/utils/handler"
)

// InitializeRedisPool 初始化redis连接池
func InitializeRedisPool(config config.Redis) {
	LinkParams := handler.RedisConfig{
		Host:       config.Host,
		Port:       config.Port,
		Db:         config.Db,
		Password:   config.Password,
		MaxRetries: config.MaxRetries,
	}
	handler.InitializeRedisPool(LinkParams)
}
