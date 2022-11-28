/**
 * Created by PhpStorm.
 * @file   token.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/28 22:46
 * @desc   token.go
 */

package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func CheckToken(context context.Context, rpcToken string) error {
	md, ok := metadata.FromIncomingContext(context)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	if token, ok := md["token"]; !ok || token[0] != rpcToken {
		return status.Error(codes.Unauthenticated, "Token认证信息无效!")
	}
	return nil
}
