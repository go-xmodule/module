/**
* Created by GoLand
* @file system.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/11 9:42 上午
* @desc system.go
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SystemMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		// t := time.Now()
		// fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		// context.Set("request", "中间件")
		// status := context.Writer.Status()
		// //fmt.Println("中间件执行完毕", status)
		// t2 := time.Since(t)
		// //fmt.Println("time:", t2)
	}
}
