/**
* Created by GoLand
* @file base_model.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/1/27 2:57 下午
* @desc base_model.go
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/server/model"
	"github.com/go-utils-module/module/utils"
	"github.com/jinzhu/gorm"
)

type Where interface{}

// Database 基础模型
type Database struct {
	db     *gorm.DB
	Result interface{}
}

// DBHandler 数据库操作句柄
var DBHandler *Database

// NewDatabase 获取新的模型
func NewDatabase(db *gorm.DB) *Database {
	database := new(Database)
	database.db = db
	database.Result = &[]map[string]interface{}{}
	return database
}

// SetModel 设置数据结果model
func (m *Database) SetModel(executeModel interface{}) *Database {
	m.Result = executeModel
	return m
}

// SetResult 设置数据结果model
func (m *Database) SetResult(resultModel interface{}) *Database {
	m.Result = resultModel
	return m
}

// DeleteByWhere 删除数据
func (m *Database) DeleteByWhere(where interface{}) error {
	err := m.db.Where(where).Delete(m.Result).Error
	if utils.HasErr(err, global.DataDeleteErr) {
		return err
	}
	return nil
}

// First 根据条件获取一条数据
func (m *Database) First(where Where, not ...Where) error {
	if len(not) > 0 {
		return m.db.Where(where).Not(not[0]).First(m.Result).Error
	} else {
		return m.db.Where(where).First(m.Result).Error
	}
}

func (m *Database) Begin() *gorm.DB {
	return m.db.Begin()
}

// Exist 检查数据是否存在
func (m *Database) Exist(where Where) (bool, error) {
	var count int
	err := m.db.Model(m.Result).Where(where).Count(&count).Error
	if utils.HasErr(err, global.DbErr) {
		return false, err
	}
	return count > 0, nil
}

// ExecuteSql 执行sql
func (m *Database) ExecuteSql(sql string, logMod ...bool) error {
	debug := true
	if len(logMod) > 0 {
		debug = logMod[0]
	}
	return m.db.LogMode(debug).Raw(sql).Scan(m.Result).Error
}

// Find 根据id获取一条数据
func (m *Database) Find(id interface{}) error {
	return m.db.Where("id=?", id).First(m.Result).Error
}

// Get 根据where条件查询数据
func (m *Database) Get(where ...map[string]interface{}) error {
	if len(where) > 0 {
		return m.db.Where(where[0]).Find(m.Result).Error
	} else {
		return m.db.Find(m.Result).Error
	}
}

// GetDb 获取当前数据库连接
func (m *Database) GetDb() *gorm.DB {
	return m.db
}

// WhereClosure 设置查询的闭包方法
func (m *Database) WhereClosure(where models.WhereClosure) *Database {
	m.db = where(m.db)
	return m
}

// SetWhere 设置当前查询条件
func (m *Database) SetWhere(whereMap map[string]interface{}) *Database {
	m.db = m.db.Where(whereMap)
	return m
}

// Select 设置查询的字段
func (m *Database) Select(query interface{}, args ...interface{}) *Database {
	m.db = m.db.Select(query, args)
	return m
}

// Create 创建数据
func (m *Database) Create(model models.ModelAction) error {
	return m.db.Create(model).Error
}

// Save 保存修改的数据
func (m *Database) Save(model interface{}) error {
	return m.db.Save(model).Error
}

// Delete 删除数据
func (m *Database) Delete(model interface{}) error {
	return m.db.Delete(model).Error
}

// GetByPage 分页查询数据
func (m *Database) GetByPage(pagination models.PaginationQuery) utils.PageData {
	var count int
	m.db.Model(m.Result).Count(&count)
	offset := pagination.PageSize * (pagination.PageNum - 1)
	m.db.Order(fmt.Sprintf("%s %s ", pagination.OrderBy, pagination.Order)).Offset(offset).Limit(pagination.PageSize).Find(m.Result)
	paginator := new(utils.PageUtil).Paginator(pagination.PageNum, pagination.PageSize, int64(count))
	paginator.Offset = offset
	paginator.EndOffset = pagination.PageSize + offset
	data, _ := json.Marshal(m.Result)
	mapRes := make([]map[string]interface{}, 0)
	_ = json.Unmarshal(data, &mapRes)
	return utils.PageData{
		PageInfo: paginator,
		DataList: mapRes,
	}
}
func (m *Database) SetMode() error {
	sql := "set sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';"
	return m.ExecuteSql(sql)
}
