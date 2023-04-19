/**
 * Created by Goland.
 * @file   sentry.go
 * @author 李锦 <Ljin@cavemanstudio.net>
 * @date   2023/4/7 21:48
 * @desc   sentry.go
 */

package middlewares

import (
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// SentryInit sentry初始化注册
func SentryInit() gin.HandlerFunc {
	return sentryGin.New(sentryGin.Options{
		Repanic: true,
	})
}

// SentryRegister sentry注册
func SentryRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if hub := sentryGin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	}
}
