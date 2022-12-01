/**
 * Created by GoLand
 * @file   mysql.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 16:08
 * @desc   mysql.go
 */

package system

import (
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/module/utils"
	"github.com/go-utils-module/module/utils/dirver"
	"github.com/go-utils-module/module/utils/handler"
	"github.com/jinzhu/gorm"
)

// InitializeDatabase 初始化数据库连接
func InitializeDatabase(conf config.Database) *gorm.DB {
	db, err := dirver.InitializeConsoleDB(dirver.LinkParams{
		Host:        conf.Database.Host,
		Port:        conf.Database.Port,
		UserName:    conf.Database.UserName,
		Password:    conf.Database.Password,
		DbName:      conf.Database.DbName,
		MaxOpenConn: conf.MaxOpenConn,
		MaxIdleConn: conf.MaxIdleConn,
		Mode:        conf.Database.Mode,
	})
	if err != nil {
		utils.Logger.Fatalln("初始化系统-连接管理后台数据库异常。", err)
	}
	handler.DBHandler = handler.NewDatabase(db)
	return db
}
