/**
 * Created by GoLand
 * @file   errors.go
* @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/5/26 19:36
 * @desc   错误信息定义,沿用nakama，所有的错误码都是201 ，故此系统定义的错误码都是201,无需单独定义
*/

package global

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

// 系统功能asdfasdf
const (
	StartServerErr           ErrCode = 101000 + iota // 启动服务异常
	SystemErr                                        // 系统异常
	SystemInitFail                                   // 系统初始化失败
	listenConfigErr                                  // 配置文件监控失败
	ParamsError                                      // 参数异常，请检查
	ConnectMysqlErr                                  // 连接数据库异常
	RequestOvertimeErr                               // 请求发起时间超时
	SignErr                                          // 参数签名异常
	GetNoticeConfigErr                               // 获取系统通知配置异常
	GetGameConfigErr                                 // 获取游戏配置异常
	GetChannelConfigErr                              // 获取发布频道配置异常
	GetLogConfigErr                                  // 获取日志配置
	GetApiConfigErr                                  // 获取Api配置
	GetDbConfigErr                                   // 获取数据库配置异常
	GetGRPCConfigErr                                 // 获取GRPC配置异常
	GetSystemConfigErr                               // 获取系统配置异常
	RedisPushErr                                     // Redis push 数据异常
	RedisPublishErr                                  // Redis 发布消息异常
	NeTRequestErr                                    // 网络请求异常
	DataSaveErr                                      // DB数据编辑异常
	DataAddErr                                       // DB数据添加异常
	DataGetErr                                       // DB数据获取异常
	GetNakamaConfigErr                               // 获取Nakama配置异常
	PublishDataErr                                   // 数据发布异常
	DbErr                                            // 数据库异常
	DataDeleteErr                                    // DB数据删除异常
	NoTokenErr                                       // 无Token认证信息
	TokenErr                                         // Token认证信息无效
	GetTokenErr                                      // 获取Token信息异常
	GetLeaderboardListErr                            // 获取Nakama排行榜数据列表异常
	GetLeaderboardDetailErr                          // 获取Nakama排行榜数据详情异常
	ParseJsonDataErr                                 // 解析Nakama json数据异常
	GetAccountListErr                                // 获取Nakama账户列表异常
	DeleteAccountErr                                 // 删除Nakama账户列表异常
	EditeAccountErr                                  // 编辑Nakama账户列表异常
	GetAccountDetailErr                              // 获取Nakama账户详情异常
	GetAccountBanListErr                             // 获取Nakama禁用账户列表异常
	DeleteLeaderboardErr                             // 删除Nakama排行榜数据异常
	AccountUnlinkErr                                 // 删除Nakama账户好友关联异常
	GetAccountFriendErr                              // 获取Nakama账户好友异常
	DeleteAccountFriendErr                           // 删除Nakama账户好友异常
	AccountEnableErr                                 // 启用Nakama账户异常
	AccountDisableErr                                // 禁用Nakama账户异常
	GetMatchDataErr                                  // 获取Nakama比赛数据异常
	GetMatchStateErr                                 // 获取Nakama比赛状态数据异常
	AccountLoginErr                                  // Nakama账户登录异常
	AccountTokenExpressErr                           // Nakama Token过期异常
	GetGameDataErr                                   // 获取Nakama数据异常
	ExecuteAfterDeleteFunErr                         // 执行删除后方法异常
	ExecuteAfterEditFunErr                           // 执行编辑后方法异常
	ExecuteBeforeEditFunErr                          // 执行编辑前方法异常
	ExecuteBeforeAddFunErr                           // 执行添加前方法异常
	LinkMysqlErr                                     // 连接数据库异常
	GetSingleDataErr                                 // 获取单条数据异常
	CreateUploadFileDirErr                           // 创建上传目录异常
	SystemError                                      // 系统异常，请稍后重试
	ParamsEmptyError                                 // 参数不可空，请检查
	ParamsFormatError                                // 参数格式错误，请检查
	RepeatRequestError                               // 重复请求
	InitSessionRedisErr                              // 初始化sessionRedis连接异常
	GetSystemNoticeConfigErr                         // 获取系统通知配置文件异常
	RegisterServerErr                                // 服务注册异常
	GetServerErr                                     // 获取服务异常
	GetConfigErr                                     // 获取配置异常
	ListenConfigErr                                  // 监听配置异常
	GetNamingClientErr                               // 获取服务实例异常
	GetConfigClientErr                               // 获取配置实例异常
	GetInstanceErr                                   // 获取服务实例异常
	RunModeErr                                       // 运行模式异常

)
