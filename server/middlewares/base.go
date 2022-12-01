/**
 * Created by PhpStorm.
 * @file   base.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:36
 * @desc   base.go
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-utils-module/module/global"
	"strings"
)

type BaseMiddleware struct {
}

// GetBaseUri 获取基础uri
func (m *BaseMiddleware) getBaseUri(context *gin.Context, domain string) string {
	temp := strings.Split(context.Request.RequestURI, domain)
	return temp[0]
}

func (m *BaseMiddleware) isApi(path string) bool {
	return strings.Contains(path, global.ApiV1)
}
