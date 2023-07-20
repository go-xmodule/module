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
	"github.com/golang-module/carbon"
	"github.com/x-module/module/config"
	"github.com/x-module/module/global"
	utils2 "github.com/x-module/utils/utils"
	"github.com/x-module/utils/utils/request"
	utils "github.com/x-module/utils/utils/response"
	"github.com/x-module/utils/utils/xlog"
	"strconv"
)

type ApiMiddleware struct {
	BaseMiddleware
	apiConfig config.Api
}

func NewApiMiddleware(apiConfig config.Api) *ApiMiddleware {
	middle := new(ApiMiddleware)
	middle.apiConfig = apiConfig
	utils2.JsonDisplay(middle.apiConfig)
	return middle
}

func (a *ApiMiddleware) Middleware() gin.HandlerFunc {
	return a.checkSign
}

// 检查参数签名
func (a *ApiMiddleware) checkSign(context *gin.Context) {
	ts := context.Query("ts")
	sign := context.Query("sign")
	if ts == "" || sign == "" {
		xlog.Warn(global.NoSignParamsErr.String())
		utils.ApiResponse(context, global.NoSignParamsErr)
		context.Abort()
		return
	}
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		xlog.Warn(global.NoSignParamsErr.String())
		utils.ApiResponse(context, global.NoSignParamsErr)
		context.Abort()
		return
	}
	// 接口请求超时超过系统超时
	if carbon.Now().Timestamp()-timestamp > a.apiConfig.Overtime {
		xlog.Warn(global.RequestOvertimeErr.String())
		utils.ApiResponse(context, global.RequestOvertimeErr)
		context.Abort()
		return
	}
	newSign := request.RequestSign(ts, a.apiConfig.Secret)
	if newSign != sign {
		xlog.Warn(global.SignErr.String())
		utils.ApiResponse(context, global.SignErr)
		context.Abort()
		return
	}
}
