/**
 * Created by Goland.
 * @file   player.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/8/2 20:11
 * @desc   player.go
 */

package service

//func GetPlayer() {
//	// 调用其他服务-玩家服务-获取玩家信息
//	player, err := player.NewPlayerService(PlayerServiceName, GrpcService.Client()).GetPlayer(ctx, &player2.GetPlayerRequest{
//		PlayerId: request.PlayerId,
//	})
//	if err == gorm.ErrRecordNotFound {
//		xlog.Warn(global.GetPlayerErr.String())
//		return fmt.Errorf(global.GetPlayerErr.String())
//	} else if err != nil {
//		xlog.WithField(global2.ErrField, err).Error(global.GetPlayerErr.String())
//		return err
//	}
//}
