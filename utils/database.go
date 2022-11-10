/**
 * Created by GoLand
 * @file   sql.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/19 21:44
 * @desc   sql.go
 */

package utils

import (
	"fmt"
	"reflect"
	"strings"
)

const BaseParamsStruct = "BaseValidationParams"

type Field struct {
	Name    string
	Field   string
	Where   string
	Default string
	Value   interface{}
	Option  string
}
type SqlUtil struct {
}

// TransWhere 结构体转换查询sql
func TransWhere(params interface{}) (string, string) {
	typeOfCat := reflect.TypeOf(params)
	valueOfCat := reflect.ValueOf(params)
	search := map[string]Field{}
	res := getValue(typeOfCat, valueOfCat, search)
	if search["StatisticalType"].Value == nil {
		return "", formatSql(res)
	}
	return search["StatisticalType"].Value.(string), formatSql(res)
}

// 获取Struct的值
func getValue(typeOfCat reflect.Type, valueOfCat reflect.Value, search map[string]Field) map[string]Field {
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		if fieldType.Name == BaseParamsStruct {
			getValue(valueOfCat.Field(i).Type(), valueOfCat.Field(i), search)
		} else {
			if reflect.Slice == valueOfCat.Field(i).Kind() {
				var value []string
				for j := 0; j < valueOfCat.Field(i).Len(); j++ {
					value = append(value, valueOfCat.Field(i).Index(j).String())
				}
				search[fieldType.Name] = Field{
					Name:    fieldType.Name,
					Field:   fieldType.Tag.Get("field"),
					Where:   fieldType.Tag.Get("where"),
					Default: fieldType.Tag.Get("default"),
					Option:  fieldType.Tag.Get("option"),
					Value:   value,
				}
			} else {
				search[fieldType.Name] = Field{
					Name:    fieldType.Name,
					Field:   fieldType.Tag.Get("field"),
					Where:   fieldType.Tag.Get("where"),
					Default: fieldType.Tag.Get("default"),
					Option:  fieldType.Tag.Get("option"),
					Value:   valueOfCat.Field(i).String(),
				}
			}
		}
	}
	return search
}

// 转换sql
func formatSql(params map[string]Field) string {
	sql := " 1 "
	for _, v := range params {
		switch v.Value.(type) {
		case string:
			value := v.Value.(string)
			if value == "" && v.Default != "" {
				value = v.Default
			}
			if v.Where != "0" && value != "" {
				if v.Name == "StartDate" {
					value += " 00:00:00"
				}
				if v.Name == "EndDate" {
					value += " 23:59:59"
				}
				option := "="
				if v.Option != "" {
					option = v.Option
				}
				sql += fmt.Sprintf(" and %s %s '%s'", v.Field, option, value)
			}
		case []string:
			value := strings.Join(v.Value.([]string), "','")
			option := " in "
			if value == "" && v.Default != "" {
				value = v.Default
				if v.Option != "" {
					option = v.Option
				}
			}
			if v.Where != "0" && value != "" {
				if v.Where != "0" && v.Value != "" {
					sql += fmt.Sprintf(" and %s %s('%s')", v.Field, option, value)
				}
			}
		}
	}

	return sql
}
