/**
 * Created by GoLand
 * @file   mysql.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 16:08
 * @desc   mysql.go
 */

package system

import (
	"github.com/x-module/module/config"
	"github.com/x-module/utils/dirver"
	"github.com/x-module/utils/handler"
)

// InitializeDatabase 初始化数据库连接
func InitializeDatabase(conf config.Database) {
	LinkParams := dirver.LinkParams{
		Host:        conf.Database.Host,
		Port:        conf.Database.Port,
		UserName:    conf.Database.UserName,
		Password:    conf.Database.Password,
		DbName:      conf.Database.DbName,
		MaxOpenConn: conf.MaxOpenConn,
		MaxIdleConn: conf.MaxIdleConn,
		Mode:        conf.Database.Mode,
	}
	handler.InitializeMysql(LinkParams)
}
