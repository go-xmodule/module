/**
 * Created by PhpStorm.
 * @file   mysql.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:06
 * @desc   mysql.go
 */

package config

import (
	"github.com/go-xmodule/module/global"
	utils "github.com/go-xmodule/utils/utils/config"
	"log"
)

type Database struct {
	Type        string `yaml:"type"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	Database    Connect
}
type Connect struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"database"`
	UserName string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"timeZone"`
	Charset  string `yaml:"charset"`
	Mode     string `yaml:"mode"`
}

// DatabaseConfigFile Database配置文件
const DatabaseConfigFile = "database.yaml"

// InitDatabaseConfig Game统配置
func InitDatabaseConfig(config any) {
	path := utils.GetConfigFile(DatabaseConfigFile)
	err := utils.GetConfig(path, config)
	if err != nil {
		log.Fatal(err, global.GetDbConfigErr.String())
	}
}
