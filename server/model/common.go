/**
 * Created by PhpStorm.
 * @file   base.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/9 19:33
 * @desc   base.go
 */

package models

import (
	"time"
)

// CommonModel 数据库类型基类
type CommonModel struct {
	Id        int       `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
