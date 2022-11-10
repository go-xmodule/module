/**
 * Created by GoLand
 * @file   http.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 14:51
 * @desc   http.go
 */

package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// IsStaticRequest 判断是否是静态文件请求
func IsStaticRequest(context *gin.Context) bool {
	if strings.Contains(context.Request.URL.Path, "/image/upload/") ||
		strings.Contains(context.Request.URL.Path, "/admin/") ||
		strings.Contains(context.Request.URL.Path, "/favicon.ico") ||
		strings.Contains(context.Request.URL.Path, ".js") {
		return true
	} else {
		return false
	}
}
