/**
 * Created by GoLand
 * @file   common.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/15 16:34
 * @desc   数据验证和绑定基类
 */

package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-utils-module/module/utils"
)

// BaseValidation 基类
type BaseValidation struct {
}

// BaseValidationParams 基础参数
type BaseValidationParams struct {
}

// ParamsValidationInter 验证接口
type ParamsValidationInter interface {
	// ParamsValidation 参数绑定
	ParamsValidation(context *gin.Context, paramsStruct interface{}) (interface{}, error)
}

func (i *BaseValidation) ParamsValidation(context *gin.Context, paramsStruct interface{}) (interface{}, error) {
	params, _ := context.Get("params")
	if err := utils.Validation(params.([]byte), paramsStruct); utils.CheckErr(err) {
		return params, err
	}
	return paramsStruct, nil
}
