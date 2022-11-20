// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package global

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-200]
	_ = x[StartServerErr-101000]
	_ = x[SystemErr-101001]
	_ = x[ParamsError-101002]
	_ = x[ConnectMysqlErr-101003]
	_ = x[RequestOvertimeErr-101004]
	_ = x[SignErr-101005]
	_ = x[GetChannelConfigErr-101006]
	_ = x[GetLogConfigErr-101007]
	_ = x[GetApiConfigErr-101008]
	_ = x[GetDbConfigErr-101009]
	_ = x[GetSystemConfigErr-101010]
	_ = x[RedisPushErr-101011]
	_ = x[RedisPublishErr-101012]
	_ = x[NeTRequestErr-101013]
	_ = x[DataSaveErr-101014]
	_ = x[DataAddErr-101015]
	_ = x[DataGetErr-101016]
	_ = x[GetNakamaConfigErr-101017]
	_ = x[PublishDataErr-101018]
	_ = x[DbErr-101019]
	_ = x[DataDeleteErr-101020]
	_ = x[GetTokenErr-101021]
	_ = x[GetLeaderboardListErr-101022]
	_ = x[GetLeaderboardDetailErr-101023]
	_ = x[ParseJsonDataErr-101024]
	_ = x[GetAccountListErr-101025]
	_ = x[DeleteAccountErr-101026]
	_ = x[EditeAccountErr-101027]
	_ = x[GetAccountDetailErr-101028]
	_ = x[GetAccountBanListErr-101029]
	_ = x[DeleteLeaderboardErr-101030]
	_ = x[AccountUnlinkErr-101031]
	_ = x[GetAccountFriendErr-101032]
	_ = x[DeleteAccountFriendErr-101033]
	_ = x[AccountEnableErr-101034]
	_ = x[AccountDisableErr-101035]
	_ = x[GetMatchDataErr-101036]
	_ = x[GetMatchStateErr-101037]
	_ = x[AccountLoginErr-101038]
	_ = x[AccountTokenExpressErr-101039]
	_ = x[GetGameDataErr-101040]
}

const (
	_ErrCode_name_0 = "Success"
	_ErrCode_name_1 = "启动服务异常系统异常参数异常，请检查连接数据库异常请求发起时间超时参数签名异常获取发布频道配置异常获取日志配置获取Api配置获取数据库配置异常获取系统配置异常Redis push 数据异常Redis 发布消息异常网络请求失败DB数据编辑异常DB数据添加异常DB数据获取异常获取Nakama配置异常数据发布失败数据库异常DB数据删除异常获取Token信息异常获取Nakama排行榜数据列表异常获取Nakama排行榜数据详情异常解析Nakama json数据异常获取Nakama账户列表异常删除Nakama账户列表异常编辑Nakama账户列表异常获取Nakama账户详情异常获取Nakama禁用账户列表异常删除Nakama排行榜数据异常删除Nakama账户好友关联异常获取Nakama账户好友异常删除Nakama账户好友异常启用Nakama账户异常禁用Nakama账户异常获取Nakama比赛数据异常获取Nakama比赛状态数据异常Nakama账户登录异常Nakama Token过期异常获取Nakama数据异常"
)

var (
	_ErrCode_index_1 = [...]uint16{0, 18, 30, 54, 75, 99, 117, 147, 165, 180, 207, 231, 254, 278, 296, 316, 336, 356, 380, 398, 413, 433, 456, 495, 534, 563, 593, 623, 653, 683, 719, 752, 788, 818, 848, 872, 896, 926, 962, 986, 1010, 1034}
)

func (i ErrCode) String() string {
	switch {
	case i == 200:
		return _ErrCode_name_0
	case 101000 <= i && i <= 101040:
		i -= 101000
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}