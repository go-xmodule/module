/**
* Created by GoLand
* @file main.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/1/26 6:19 下午
* @desc   主程序入口
 */

package global

import "embed"

func BuildServer(template embed.FS, assets embed.FS) {
	TemplatePath = template
	AssetsPath = assets
}
