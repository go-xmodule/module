/**
 * Created by PhpStorm.
 * @file   ap.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:34
 * @desc   ap.go
 */

package config

// Api 接口配置
type Api struct {
	Secret   string `yaml:"secret"`
	Overtime int64  `yaml:"overtime"`
}
