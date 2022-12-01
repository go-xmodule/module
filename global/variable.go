/**
 * Created by GoLand
 * @file   variable.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 22:16
 * @desc   variable.go
 */

package global

import "embed"

var TemplatePath embed.FS
var AssetsPath embed.FS

// 运行环境定义
const (
	// DevMode 开发环境
	DevMode = "dev"
	// TestMode 测试环境
	TestMode = "test"
	// OnlineMode 生成环境
	OnlineMode = "online"
)

// 常量定义
const (
	ApiV1                = "/api/v1"
	RequestParams string = "params"
)

const (
	ErrField = "error"
)
