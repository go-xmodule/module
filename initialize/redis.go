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
	"github.com/go-utils-module/module/dirver"
)

// InitializeRedisPool 初始化redis连接池
func InitializeRedisPool(config config.Redis) dirver.RedisClient {
	c, err := dirver.NewRedis().Connect(config.Host, config.Port, config.Password, config.Db)
	if err != nil {
		panic(err)
		return dirver.RedisClient{}
	}
	return *c
}
