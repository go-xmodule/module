/**
 * Created by PhpStorm.
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:06
 * @desc   redis.go
 */

package config

// Redis 配置
type Redis struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Db         int    `yaml:"db"`
	Password   string `yaml:"password"`
	MaxRetries int    `yaml:"maxRetries"`
}
