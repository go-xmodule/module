/**
 * Created by PhpStorm.
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/9 20:41
 * @desc   redis.go
 */

package subscribe

import (
	"github.com/go-utils-module/module/utils"
	"github.com/go-utils-module/module/utils/dirver"
	"github.com/go-utils-module/module/utils/handler"
)

type RedisSubscribe struct {
}

func NewRedisSubscribe() *RedisSubscribe {
	return new(RedisSubscribe)
}

// Subscribe 订阅消息
func (s *RedisSubscribe) Subscribe(channel string, callback SubscribeCallback) {
	utils.Logger.Debug("start subscribe data, channel:", channel)
	messageList := handler.RedisHandler.Subscribe(channel)
	for message := range messageList {
		// 处理消息
		utils.Logger.Debug("consumer data:", utils.Json(message))
		var data dirver.SubscribeData
		_ = utils.TransInterfaceToStruct(message, &data)
		callback(data.Payload)
	}
}

// Publish 发布数据
func (s *RedisSubscribe) Publish(channel string, message interface{}) error {
	utils.Logger.Debug("start publish data, message:", utils.Json(message))
	_, err := handler.RedisHandler.Publish(channel, message)
	if err != nil {
		utils.Logger.Error(PublishErr, err.Error())
		return err
	}
	utils.Logger.Debug("publish data success")
	return nil
}
