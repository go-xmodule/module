/**
* Created by GoLand
* @file mysql.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/1/27 11:42 上午
* @desc 初始化管理后台数据库
 */

package dirver

import (
	"fmt"
	"github.com/go-utils-module/module/utils"
	"github.com/go-utils-module/module/utils/code"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	DbType    = "mysql"
)

type LinkParams struct {
	Host        string
	Port        int
	UserName    string
	DbName      string
	Password    string
	MaxOpenConn int
	MaxIdleConn int
	Model       string
}

// InitializeConsoleDB 初始化管理后台数据库
func InitializeConsoleDB(params LinkParams) (*gorm.DB, error) {
	linkParams := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	linkAddress := fmt.Sprintf(linkParams, params.UserName, params.Password, params.Host, params.Port, params.DbName)
	db, err := gorm.Open(DbType, linkAddress)
	if utils.HasErr(err, code.ConnectMysqlErr) {
		return nil, err
	}
	// 链接池设置
	db.DB().SetMaxOpenConns(params.MaxOpenConn)
	db.DB().SetMaxIdleConns(params.MaxIdleConn)
	db.LogMode(params.Model == DebugMode)
	// db.LogMode(false)
	return db, nil
}
