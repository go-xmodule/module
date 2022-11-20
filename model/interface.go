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

// PaginationQuery 分页查询
type PaginationQuery struct {
	PageSize int
	PageNum  int
	// OrderBy 小写的字段名称
	OrderBy string
	// Order 默认是'desc', 可选的: 'desc', 'asc'
	Order string
}

// ActonInfo 当前数据模型路由介绍
type ActonInfo struct {
	Action     string
	ParentMenu string
	ActionType string
	ActionTag  string
}
type WhereClosure func(*gorm.DB) *gorm.DB

type ModelAction interface {
	ModelInfo() ActonInfo
	DataId() int
	TableName() string
	DbConfig() ConfigType
}
