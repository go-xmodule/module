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

// 运行模式定义
const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

// 常量定义
const (
	ApiV1                = "/api/v1"
	RequestParams string = "params"
)

const (
	ErrField = "error"
)
