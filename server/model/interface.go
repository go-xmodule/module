/**
 * Created by GoLand
 * @file   interface.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/7/8 17:50
 * @desc   interface.go
 */

package models

import "github.com/jinzhu/gorm"

// ConfigType 数据库类型
type ConfigType int

type WhereClosure func(*gorm.DB) *gorm.DB

// ActonInfo 当前数据模型路由介绍
type ActonInfo struct {
	Action     string
	ParentMenu string
	ActionType string
	ActionTag  string
}

type ModelAction interface {
	ModelInfo() ActonInfo
	DataId() int
	TableName() string
	// DbConfig() ConfigType
}
