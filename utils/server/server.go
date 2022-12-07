/**
 * Created by goland.
 * @file   server.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/7 14:48
 * @desc   server.go
 */

package server

import (
	"errors"
	"fmt"
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"log"
	"strings"
)

type ServerInfo struct {
	Describe string `json:"describe,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
}

var ServerList = map[string]ServerInfo{}

type ServerGroup struct {
	ServiceName string   `param:"serviceName"` // required
	Clusters    []string `param:"clusters"`    // optional,default:DEFAULT
	GroupName   string   `param:"groupName"`   // optional,default:DEFAULT_GROUP
}

// SubscribeServer 订阅服务
func SubscribeServer(client naming_client.INamingClient, serverList []ServerGroup) {
	for _, server := range serverList {
		err := nacos.SubscribeServer(nacos.SubscribeServerParams{
			Client:      client,
			ServiceName: server.ServiceName,
			GroupName:   server.GroupName,
			SubscribeCallback: func(services []model.Instance, err error) {
				utils.Logger.Debugf("server config change:%+v", services)
				serviceName := strings.Split(services[0].ServiceName, "@@")
				var serverInfo ServerInfo
				_ = utils.TransInterfaceToStruct(services[0].Metadata, &serverInfo)
				ServerList[serviceName[0]] = serverInfo
			},
		})
		if err != nil {
			log.Fatal(global.SystemInitFail.String())
		}
	}
}

// InitServer 当前服务配置初始化
func InitServer(client naming_client.INamingClient, serverList []ServerGroup) {
	for _, server := range serverList {
		service, err := nacos.GetService(nacos.GetServerParams{
			Client:      client,
			ServiceName: server.ServiceName,
			GroupName:   server.GroupName,
		})
		if err != nil {
			log.Fatal(global.SystemInitFail.String())
		}
		if len(service.Hosts) > 0 {
			serviceName := strings.Split(service.Hosts[0].ServiceName, "@@")
			var serverInfo ServerInfo
			_ = utils.TransInterfaceToStruct(service.Hosts[0].Metadata, &serverInfo)
			ServerList[serviceName[0]] = serverInfo
		}
	}
}

// GetServerAddress 获取服务信息
func GetServerAddress(serverName string) (string, error) {
	if server, ok := ServerList[serverName]; !ok {
		utils.Logger.Errorf("%s,server:%s", global.UnknownServerErr.String(), serverName)
		return "", errors.New(global.UnknownServerErr.String())
	} else {
		return fmt.Sprintf("%s:%s", server.Host, server.Port), nil
	}
}
