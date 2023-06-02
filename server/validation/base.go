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
	"github.com/x-module/module/global"
	"github.com/x-module/module/utils"
	"github.com/x-module/utils/utils/xlog"
)

// BaseValidation 基类
type BaseValidation[T any] struct {
	context      *gin.Context
	paramsStruct T
}

func NewBaseValidation[T any](ctx *gin.Context, params T) *BaseValidation[T] {
	return &BaseValidation[T]{
		context:      ctx,
		paramsStruct: params,
	}
}

// BaseValidationParams 基础参数
type BaseValidationParams struct {
}

// ParamsValidationInter 验证接口
type ParamsValidationInter[T any] interface {
	// Validation ParamsValidation 参数绑定
	Validation(context *gin.Context) (T, error)
}

func (i *BaseValidation[T]) Validation() (T, error) {
	params, _ := i.context.Get(global.RequestParams)

	if err := utils.Validation(params.([]byte), i.paramsStruct); err != nil {
		xlog.Logger.Warn(err)
		return i.paramsStruct, err
	}
	return i.paramsStruct, nil
}
