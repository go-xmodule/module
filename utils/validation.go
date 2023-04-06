/**
 * Created by PhpStorm.
 * @file   validation.go
 * @author 李锦 <Ljin@cavemanstudio.net>
 * @date   2022/10/13 14:30
 * @desc   validation.go
 */

package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Validation(requestParams []byte, obj any) error {
	err := json.Unmarshal(requestParams, &obj)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(obj).Elem()
	t1 := reflect.ValueOf(obj)
	var checkErrors []string
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("binding") == "required" && t1.Elem().Field(i).String() == "" {
			errMsg := t.Field(i).Tag.Get("msg")
			if errMsg == "" {
				errMsg = "不能为空！"
			}
			checkErrors = append(checkErrors, fmt.Sprintf(" %s:%s", t.Field(i).Name, errMsg))
		}
	}
	if len(checkErrors) == 0 {
		return nil
	} else {
		return errors.New(strings.Join(checkErrors, ","))
	}
}
func Validation1(obj any) error {
	t := reflect.TypeOf(obj).Elem()
	t1 := reflect.ValueOf(obj)
	err := ""
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("binding") == "required" && t1.Elem().Field(i).String() == "" {
			errMsg := t.Field(i).Tag.Get("msg")
			if errMsg == "" {
				errMsg = t.Field(i).Name + " 为空！"
			}
			err += fmt.Sprintf(" %s:%s", t.Field(i).Name, errMsg)
		}
	}
	if len(err) == 0 {
		return nil
	} else {
		return errors.New(err)
	}
}
