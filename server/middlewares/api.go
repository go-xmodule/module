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
)

type ApiMiddleware struct {
	BaseMiddleware
	serverConfig config.Server
	apiConfig    config.Api
}

func NewApiMiddleware(serverConfig config.Server, apiConfig config.Api) *ApiMiddleware {
	middle := new(ApiMiddleware)
	middle.serverConfig = serverConfig
	middle.apiConfig = apiConfig
	return middle
}

func (a *ApiMiddleware) Middleware() gin.HandlerFunc {
	return a.checkSign
}

// 检查参数签名
func (a *ApiMiddleware) checkSign(context *gin.Context) {
	data, _ := context.GetRawData()
	context.Set(global.RequestParams, data)
	path := a.getBaseUri(context, a.serverConfig.Domain)
	if a.isApi(path) { // 不是api 请求
		// todo
	}
}
