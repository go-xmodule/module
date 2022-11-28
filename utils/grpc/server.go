package grpc

import (
	"google.golang.org/grpc"
	"net"
)

const (
	TCP = "tcp"
	UDP = "udp"
)

var GrpcServer = new(Server)

type ServerRegister func(server *grpc.Server)

type Server struct {
	network  string
	address  string
	register ServerRegister
}

func NewServer() *Server {
	return new(Server)
}

// Init 初始化
func (g *Server) Init(network, address string) *Server {
	g.network = network
	g.address = address
	return g
}
func (g *Server) RegisterServer(register ServerRegister) *Server {
	g.register = register
	return g
}

// Start 启动grpc 服务
func (g *Server) Start() error {
	listen, err := net.Listen(g.network, g.address)
	if err != nil {
		return err
	}
	// 实例化grpc Server
	s := grpc.NewServer()
	// 注册服务
	g.register(s)
	return s.Serve(listen)
}
