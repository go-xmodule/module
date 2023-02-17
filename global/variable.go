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
	// ================================
	DefaultNacosGroup = "DEFAULT_GROUP"
)

// 常量定义
const (
	ApiV1                = "/api/v1"
	RequestParams string = "params"
	ENVIRONMENT   string = "ENVIRONMENT"
)

const (
	ErrField = "error"
)

// 配置定义
const (
	// DatabaseConfigDataId 数据库配置
	DatabaseConfigDataId = "database"
	// SystemConfigDataId 系统配置
	SystemConfigDataId = "system"
	// NakamaConfigDataId nakama 配置
	NakamaConfigDataId = "nakama"
	// GameConfigDataId 游戏数据配置
	GameConfigDataId = "game"
	// ChannelConfigDataId 订阅评定配置
	ChannelConfigDataId = "channel"
	// NoticeConfigDataId 通知配置
	NoticeConfigDataId = "notice"
	// GRPCConfigDataId GRPC配置
	GRPCConfigDataId = "grpc"
	// ApiConfigDataId Api 配置
	ApiConfigDataId = "api"
)

type Server struct {
	Name  string `yaml:"name"`
	Group string `yaml:"group"`
	Desc  string `yaml:"desc"`
}

var ServerList = []Server{
	{
		Name:  "notification",
		Group: "notification",
		Desc:  "消息通知服务",
	},
	{
		Name:  "collection",
		Group: "collection",
		Desc:  "数据收集服务",
	},
	{
		Name:  "web",
		Group: "web",
		Desc:  "主站服务",
	},
	{
		Name:  "console",
		Group: "console",
		Desc:  "控制台服务",
	},
	{
		Name:  "analysis",
		Group: "analysis",
		Desc:  "数据分析服务",
	},
}
