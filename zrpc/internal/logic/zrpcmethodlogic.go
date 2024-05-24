package logic

import (
	"context"

	"go-zero-aula-1/common/pb"
	"go-zero-aula-1/zrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZrpcMethodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewZrpcMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZrpcMethodLogic {
	return &ZrpcMethodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ZrpcMethodLogic) ZrpcMethod(in *__.ZprcRequest) (*__.RealizarPedidoResponse, error) {
	return &__.RealizarPedidoResponse{ZrpcResponse: "Resposta do Zrpc!"}, nil
}
