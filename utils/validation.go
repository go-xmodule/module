/**
 * Created by PhpStorm.
 * @file   validation.go
 * @author 李锦 <Ljin@cavemanstudio.net>
 * @date   2022/10/13 14:30
 * @desc   validation.go
 */

package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strings"
)

// BaseValidation 基类
type BaseValidation[T any] struct {
	validate T
	ctx      context.Context
}

func NewBaseValidation[T any](ctx context.Context, validate T) *BaseValidation[T] {
	return &BaseValidation[T]{
		validate: validate,
		ctx:      ctx,
	}
}

// BaseValidationParams 基础参数
type BaseValidationParams struct {
}

func (i *BaseValidation[T]) Validation() error {
	params, _ := json.Marshal(i.validate)
	if err := Validation(params, i.validate); err != nil {
		fmt.Println("=====================================================")
		fmt.Println("request:", string(params))
		fmt.Println("=====================================================")
		Logger(i.ctx).Warn(err)
		return err
	}
	return nil
}

func Validation(requestParams []byte, obj any) error {
	err := jsoniter.UnmarshalFromString(string(requestParams), &obj)
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
	// 判断obj中时是否存在为空的字段且字段设置了default值，如果存在则将default值赋值给obj
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("binding") == "required" && t1.Elem().Field(i).String() == "" && t.Field(i).Tag.Get("default") != "" {
			t1.Elem().Field(i).SetString(t.Field(i).Tag.Get("default"))
		}
	}
	if len(checkErrors) == 0 {
		return nil
	} else {
		return errors.New(strings.Join(checkErrors, ","))
	}
}
