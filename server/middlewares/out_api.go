/**
 * Created by PhpStorm.
 * @file   api_middleware.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/1 17:32
 * @desc   api_middleware.go
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils"
	"github.com/golang-module/carbon"
)

type OutApiMiddleware struct {
	BaseMiddleware
	serverConfig config.Server
	apiConfig    config.Api
}

func NewOutApiMiddleware(serverConfig config.Server, apiConfig config.Api) *OutApiMiddleware {
	middle := new(OutApiMiddleware)
	middle.serverConfig = serverConfig
	middle.apiConfig = apiConfig
	return middle
}

func (a *OutApiMiddleware) Middleware() gin.HandlerFunc {
	return a.checkSign
}

// 检查参数签名
func (a *OutApiMiddleware) checkSign(context *gin.Context) {
	data, _ := context.GetRawData()
	context.Set(global.RequestParams, data)
	path := a.getBaseUri(context, a.serverConfig.Domain)
	if a.isApi(path) { // 不是api 请求
		params := utils.TansToMap(data)
		requestTime := int64(params["timestamp"].(float64))
		// 接口请求超时超过系统超时
		if carbon.Now(carbon.PRC).Timestamp()-requestTime > a.apiConfig.Overtime {
			utils.ApiResponse(context, global.RequestOvertimeErr)
			context.Abort()
			return
		}
		requestParamsSign := params["sign"]
		delete(params, "sign")
		params["timestamp"] = int64(params["timestamp"].(float64))
		newSign := utils.Sign(params, a.apiConfig.Secret)
		if newSign != requestParamsSign {
			utils.ApiResponse(context, global.SignErr)
			context.Abort()
			return
		}
	}
}
