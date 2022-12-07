/**
 * Created by goland.
 * @file   grpc.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/1 21:52
 * @desc   grpc.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
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

func InitGrpcConfig(client config_client.IConfigClient, group string) GrpcConfig {
	var grpcConfig GrpcConfig
	getConfigParams := nacos.GetConfigParams{
		Client: client,
		DataId: global.GRPCConfigDataId,
		Group:  group,
	}
	err := nacos.GetConfig(getConfigParams, &grpcConfig)
	if err != nil {
		log.Printf("%s,err:%s", global.GetConfigErr.String(), err.Error())
		log.Fatal(global.GetGRPCConfigErr)
	}
	return grpcConfig
}
