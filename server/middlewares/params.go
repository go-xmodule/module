/**
 * Created by goland.
 * @file   params.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/7 19:08
 * @desc   params.go
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/x-module/module/global"
)

func ParamsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, _ := context.GetRawData()
		context.Set(global.RequestParams, data)
		// path := a.getBaseUri(context, a.serverConfig.Domain)
		// if a.isApi(path) { // 不是api 请求
		// 	// todo
		// }
		context.Next()
	}
}
