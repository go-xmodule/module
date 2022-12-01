/**
 * Created by goland.
 * @file   client.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/30 10:16
 * @desc   client.go
 */

package utils

import (
	"github.com/go-utils-module/module/global"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

type GetServerParams struct {
	Client      naming_client.INamingClient `json:"client,omitempty"`
	ServiceName string                      `json:"service_name,omitempty"` // required
	Cluster     []string                    `json:"cluster_name,omitempty"` // optional,default:DEFAULT
	GroupName   string                      `json:"group_name,omitempty"`   // optional,default:DEFAULT_GROUP
}
type GetInstanceParams struct {
	Client      naming_client.INamingClient `json:"client,omitempty"`
	ServiceName string                      `json:"service_name,omitempty"` // required
	Cluster     []string                    `json:"cluster_name,omitempty"` // optional,default:DEFAULT
	GroupName   string                      `json:"group_name,omitempty"`   // optional,default:DEFAULT_GROUP
}

// GetService 获取注册中信的服务
func GetService(params GetServerParams) (model.Service, error) {
	service, err := params.Client.GetService(
		vo.GetServiceParam{
			ServiceName: params.ServiceName,
			GroupName:   params.GroupName,
			Clusters:    params.Cluster,
		})
	if err != nil {
		log.Printf("%s,err:%s", global.GetServerErr.String(), err.Error())
		return model.Service{}, err
	}
	return service, nil
}

// GetInstance 获取注册中信的服务实例
func GetInstance(params GetInstanceParams) ([]model.Instance, error) {
	instances, err := params.Client.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: params.ServiceName,
		GroupName:   params.GroupName,
		Clusters:    params.Cluster,
	})
	if err != nil {
		log.Printf("%s,err:%s", global.GetInstanceErr.String(), err.Error())
		return nil, err
	}
	return instances, nil
}
