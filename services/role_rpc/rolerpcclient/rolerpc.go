// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: role_rpc.proto

package rolerpcclient

import (
	"context"

	"role_rpc/role_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = role_rpc.Request
	Response = role_rpc.Response

	RoleRpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultRoleRpc struct {
		cli zrpc.Client
	}
)

func NewRoleRpc(cli zrpc.Client) RoleRpc {
	return &defaultRoleRpc{
		cli: cli,
	}
}

func (m *defaultRoleRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := role_rpc.NewRoleRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
