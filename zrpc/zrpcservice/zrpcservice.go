// Code generated by goctl. DO NOT EDIT.
// Source: zrpc.proto

package zrpcservice

import (
	"context"

	"go-zero-aula-1/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	RealizarPedidoResponse = __.RealizarPedidoResponse
	ZprcRequest            = __.ZprcRequest

	ZrpcService interface {
		ZrpcMethod(ctx context.Context, in *ZprcRequest, opts ...grpc.CallOption) (*RealizarPedidoResponse, error)
	}

	defaultZrpcService struct {
		cli zrpc.Client
	}
)

func NewZrpcService(cli zrpc.Client) ZrpcService {
	return &defaultZrpcService{
		cli: cli,
	}
}

func (m *defaultZrpcService) ZrpcMethod(ctx context.Context, in *ZprcRequest, opts ...grpc.CallOption) (*RealizarPedidoResponse, error) {
	client := __.NewZrpcServiceClient(m.cli.Conn())
	return client.ZrpcMethod(ctx, in, opts...)
}
