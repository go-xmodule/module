/**
 * Created by GoLand
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 16:05
 * @desc   redis.go
 */

package system

import (
	"github.com/getsentry/sentry-go"
	"github.com/x-module/module/config"
	"github.com/x-module/utils/utils/xlog"
)

// InitializeSentry 初始化Sentry
func InitializeSentry(config config.SentryConfig) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Params.Dsn,
		Environment:      config.Params.Environment,
		ServerName:       config.Params.ServerName,
		Release:          config.Params.Release,
		TracesSampleRate: config.Params.TracesSampleRate,
		AttachStacktrace: config.Params.AttachStacktrace,
		Debug:            config.Params.Debug,
	})
	if err != nil {
		xlog.Logger.Fatalf("sentry.Init: %s", err)
	}
}
