/**
* Created by GoLand
* @file reflect.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/10 10:00 上午
* @desc reflect.go
 */

package utils

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	ParseTimeTemplate = "2006-01-02T15:04:05+08:00"
	TimeTemplate      = "2006-01-02 15:04:05"
	TimeTemplate2     = "20060102150405"
)

func GetStructMap(data interface{}) map[string]interface{} {
	s, _ := json.Marshal(data)
	var mapData map[string]interface{}
	_ = json.Unmarshal(s, &mapData)
	for field, value := range mapData {
		if field == "created_at" || field == "updated_at" || field == "deleted_at" {
			if value != nil {
				t, _ := time.ParseInLocation(ParseTimeTemplate, value.(string), time.Local)
				mapData[field] = t.Format(TimeTemplate)
			}
		}
	}
	return mapData
}

func SetStructField(ptr interface{}, fields map[string]string) {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		name := ToLine(fieldInfo.Name)
		if value, ok := fields[name]; ok {
			switch v.FieldByName(fieldInfo.Name).Type().Kind().String() {
			case "int":
				val, _ := strconv.Atoi(value)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(val))
			case "string":
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}
		}
	}
}

func GetStructField(ptr interface{}) []string {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	var fieldNames []string
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("json")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		name = strings.Split(name, ",")[0]
		fieldNames = append(fieldNames, name)
	}
	return fieldNames
}
