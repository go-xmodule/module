/**
 * Created by PhpStorm.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:19
 * @desc   log.go
 */

package config

// Log 日志设置
type Log struct {
	Path string `yaml:"path"`
	File string `yaml:"file"`
	Mode string `yaml:"mode"`
}
