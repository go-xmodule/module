/**
 * Created by goland.
 * @file   grpc.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/1 21:52
 * @desc   grpc.go
 */

package config

import (
	"github.com/go-xmodule/module/global"
	utils "github.com/go-xmodule/utils/utils/config"
	"log"
)

type GrpcConfig struct {
	Network Network `yaml:"params"`
	Auth    Auth    `yaml:"auth"`
}
type Auth struct {
	Token string `yaml:"token"`
}

type Network struct {
	Network string `yaml:"network"`
	Ip      string `yaml:"ip"`
	Port    int    `yaml:"port"`
}

const GRPCConfigFile = "grpc.yaml"

func InitGrpcConfig() GrpcConfig {
	var config GrpcConfig
	path := utils.GetConfigFile(GRPCConfigFile)
	err := utils.GetConfig(path, &config)
	if err != nil {
		log.Fatal(err, global.GetGRPCConfigErr.String())
		return GrpcConfig{}
	}
	return config
}
