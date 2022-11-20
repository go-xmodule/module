/**
 * Created by PhpStorm.
 * @file   system.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 02:10
 * @desc   system.go
 */

package config

// System 系统配置
type System struct {
	Mode    string `yaml:"mode"`
	TempDir string `yaml:"tempDir"`
}
