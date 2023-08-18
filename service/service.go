/**
 * Created by Goland.
 * @file   service.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/7/20 17:21
 * @desc   service.go
 */

package service

const (
	// SkinServiceName 皮肤服务名称
	SkinServiceName = "go.micro.api.skin"
	// PlayerServiceName 玩家服务名称
	PlayerServiceName = "go.micro.api.player"
	// PendantServiceName 挂件服务名称
	PendantServiceName = "go.micro.api.pendant"
	// BadgeServiceName 狗牌服务名称
	BadgeServiceName = "go.micro.api.badge"
	// BuffServiceName buff服务名称
	BuffServiceName = "go.micro.api.buff"
	// NoticeServiceName 通知服务名称
	NoticeServiceName = "go.micro.api.notice"
	// ConfigServiceName 配置服务名称
	ConfigServiceName = "go.micro.api.config"
	// RewardServiceName 玩家奖励服务
	RewardServiceName = "go.micro.api.reward"
	// PassServiceName 玩家通行证服务
	PassServiceName = "go.micro.api.pass"
	// ShopServiceName 商城服务
	ShopServiceName = "go.micro.api.shop"
	//	PropServiceName 道具服务
	PropServiceName = "go.micro.api.prop"
	// MicaServiceName mica服务
	MicaServiceName = "go.micro.api.mica"
	// TaskServiceName 任务服务
	TaskServiceName = "go.micro.api.task"
	// CollectionServiceName 收集服务
	collectionServiceName = "go.micro.api.collection"
	// QueryServiceName 队列服务
	QueryServiceName = "go.micro.api.query"
	// EventServiceName 事件服务
	EventServiceName = "go.micro.api.event"
)

// ServicesMap 服务名称映射
var ServicesMap = map[string]string{
	"skin":       SkinServiceName,
	"player":     PlayerServiceName,
	"pendant":    PendantServiceName,
	"badge":      BadgeServiceName,
	"buff":       BuffServiceName,
	"notice":     NoticeServiceName,
	"config":     ConfigServiceName,
	"reward":     RewardServiceName,
	"pass":       PassServiceName,
	"shop":       ShopServiceName,
	"prop":       PropServiceName,
	"mica":       MicaServiceName,
	"task":       TaskServiceName,
	"collection": collectionServiceName,
	"query":      QueryServiceName,
	"event":      EventServiceName,
}
