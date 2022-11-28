/**
 * Created by PhpStorm.
 * @file   client.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/27 22:51
 * @desc   client.go
 */

package grpc

import (
	"context"
	"google.golang.org/grpc"
)

var GrpcClient = new(Client)

type Client struct {
	address string
	token   string
	connect *grpc.ClientConn
}

func NewClient() *Client {
	return new(Client)
}

// Init 初始化
func (g *Client) Init(address string, token string) *Client {
	g.address = address
	g.token = token
	return g
}

func (g *Client) Close() error {
	return g.connect.Close()
}

func (g *Client) Connect() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(NewCustomCredential(g.token)))
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(g.address, opts...)
	if err != nil {
		return nil, err
	}
	g.connect = conn
	return conn, err
}

// CustomCredential 自定义认证
type CustomCredential struct {
	token string
}

func NewCustomCredential(token string) *CustomCredential {
	return &CustomCredential{
		token: token,
	}
}

// GetRequestMetadata 实现自定义认证接口
func (c CustomCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": c.token,
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c CustomCredential) RequireTransportSecurity() bool {
	return false
}
