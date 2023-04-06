/**
 * Created by goland.
 * @file   game.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/1 20:46
 * @desc   game.go
 */

package config

import (
	"github.com/go-xmodule/module/global"
	utils "github.com/go-xmodule/utils/utils/config"
	"log"
)

// GameConfigFile Game配置文件
const GameConfigFile = "game.yaml"

// InitGameConfig Game统配置
func InitGameConfig(config any) {
	path := utils.GetConfigFile(GameConfigFile)
	err := utils.GetConfig(path, config)
	if err != nil {
		log.Fatal(err, global.GetGameConfigErr.String())
	}
}
