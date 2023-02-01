/**
 * Created by PhpStorm.
 * @file   base_api.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/19 16:40
 * @desc   base_api.go
 */

package controller

type BaseApiController struct {
	BaseController
}

// Response 响应结构体
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
