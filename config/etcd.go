/**
 * Created by Goland.
 * @file   etcde.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/7/24 20:39
 * @desc   etcde.go
 */

package config

type Etcd struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
