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
	"github.com/go-xmodule/module/config"
	"github.com/go-xmodule/module/global"
	"github.com/go-xmodule/utils/utils/request"
	utils "github.com/go-xmodule/utils/utils/response"
	"github.com/go-xmodule/utils/utils/xlog"
	"github.com/golang-module/carbon"
	"strconv"
)

type ApiMiddleware struct {
	BaseMiddleware
	apiConfig config.Api
}

func NewApiMiddleware(apiConfig config.Api) *ApiMiddleware {
	middle := new(ApiMiddleware)
	middle.apiConfig = apiConfig
	return middle
}

func (a *ApiMiddleware) Middleware() gin.HandlerFunc {
	return a.checkSign
}

// 检查参数签名
func (a *ApiMiddleware) checkSign(context *gin.Context) {
	ts := context.Query("ts")
	sign := context.Query("sign")
	xlog.Logger.Debug("--------------------ts:", ts, "  sign:", sign)

	if ts == "" || sign == "" {
		xlog.Logger.Debug("--------------------ts or sign is null")
		utils.ApiResponse(context, global.NoSignParamsErr)
		context.Abort()
		return
	}
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		xlog.Logger.Debug("-------------------- timestamp error")

		utils.ApiResponse(context, global.NoSignParamsErr)
		context.Abort()
		return
	}
	// 接口请求超时超过系统超时
	if carbon.Now().Timestamp()-timestamp > a.apiConfig.Overtime {
		xlog.Logger.Debug("--------------------接口请求超时超过系统超时")

		utils.ApiResponse(context, global.RequestOvertimeErr)
		context.Abort()
		return
	}
	newSign := request.RequestSign(ts, a.apiConfig.Secret)
	if newSign != sign {
		xlog.Logger.Debug("-------------sa.apiConfig.Secret:", a.apiConfig.Secret, "  ts:", a.apiConfig.Secret, "  sign:", sign, " newSign:", newSign)
		xlog.Logger.Debug("-------------------- 签名错误")

		utils.ApiResponse(context, global.SignErr)
		context.Abort()
		return
	}

}
