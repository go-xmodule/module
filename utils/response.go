/**
 * Created by Goland.
 * @file   display.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/11 16:17
 * @desc   display.go
 */

package utils

import (
	"github.com/go-utils-module/module/utils/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 响应结构体
type ResponseData struct {
	Code code.ErrCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

// ApiResponse 异常通知
func ApiResponse(context *gin.Context, errorCode code.ErrCode, data ...interface{}) {
	response := ResponseData{
		Code: errorCode,
		Msg:  errorCode.String(),
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	context.JSON(http.StatusOK, response)
}
