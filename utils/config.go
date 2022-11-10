/**
* Created by Goland
* @file load_config.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2021/10/28 11:28 上午
* @desc
 */

package utils

import (
	"errors"
	"github.com/druidcaesa/gotool"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

const OnlineEnv = "online"
const TestEnv = "test"

// GetConfig 获取配置
func GetConfig(path string, config interface{}) error {
	exists := gotool.FileUtils.Exists(path)
	if !exists {
		return errors.New("config file:" + path + " is not found")
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if yaml.Unmarshal(content, config) != nil {
		return err
	}
	return nil
}

// GetConfigFile 获取当前运行环境下的配置
func GetConfigFile(configFile string) string {
	mode := os.Getenv("ENVIRONMENT")
	config := path.Join("config", configFile)
	if mode != "" {
		config = path.Join("config", mode, configFile)
	}
	return config
}
