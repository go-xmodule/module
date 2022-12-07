/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/1 21:00
 * @desc   utils.go
 */

package config

import (
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"log"
)

func ListenConfig(client config_client.IConfigClient, dataId string, group string, onChange nacos.OnChange) {
	// 监听数据库配置更新
	err := nacos.ListenConfig(nacos.ListenConfigParams{
		Client:   client,
		DataId:   dataId,
		Group:    group,
		OnChange: onChange,
	})
	if err != nil {
		log.Fatal(global.ListenConfigErr.String(), err.Error())
	}

}
