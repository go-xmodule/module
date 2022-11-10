/**
 * Created by GoLand
 * @file   errors.go
* @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/5/26 19:36
 * @desc   错误信息定义,沿用nakama，所有的错误码都是201 ，故此系统定义的错误码都是201,无需单独定义
*/

package code

/**
错误码规则: 错误码暂定都是5位数字
. 错误码为 0 表示成功，其他都表示错误。
. 数字 1 开头的错误码表示系统级别的错误，比如缺少某种字符集，连不上数据库之类的，系统级的错误码不需要分模块，可以按照自增方式进行添加。
. 数字 2 开头的错误码表示命令行错误
. 数字 3 开头的错误码表示内部API错误
. 数字 4 开头的错误码表示调用外部API错误
. 第二、三位标识功能
. 第四、五位、六位标识错误
. 例如：201001  内部的包列表接口参数错误
**/

type ErrCode int64

//go:generate stringer -type ErrCode -linecomment

const (
	Success ErrCode = 200 // Success
)

// 系统功能
const (
	SystemErr          ErrCode = 101000 + iota // 系统异常
	ParamsError                                // 参数异常，请检查
	ConnectMysqlErr                            // 连接数据库异常
	RequestOvertimeErr                         // 请求发起时间超时
	SignErr                                    // 参数签名异常
	NeTRequestErr                              // 网络请求失败
	ParseJsonDataErr                           // 解析json数据异常
)

// 命令行游戏服务信息相关
const (
	SaveGameInfoError      ErrCode = 401000 + iota // 保存游戏信息异常
	GetGameInfoErr                                 // 获取游戏信息异常
	GetTokenErr                                    // 获取Token信息异常
	GetGameDataErr                                 // 获取Nakama数据异常
	AccountLoginErr                                // Nakama账户登录异常
	AccountTokenExpressErr                         // Nakama Token过期异常
	SaveGameStatisticsErr                          // 保存游戏统计数据异常
)
