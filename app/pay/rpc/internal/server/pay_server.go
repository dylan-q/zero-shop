// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package server

import (
	"context"

	"zero-shop/app/pay/rpc/internal/logic"
	"zero-shop/app/pay/rpc/internal/svc"
	"zero-shop/app/pay/rpc/pb"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPayServer
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) Create(ctx context.Context, in *pb.PayReq) (*pb.PayResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *PayServer) Third(ctx context.Context, in *pb.ThirdReq) (*pb.ThirdResp, error) {
	l := logic.NewThirdLogic(ctx, s.svcCtx)
	return l.Third(in)
}

func (s *PayServer) GetThirdByOrderSn(ctx context.Context, in *pb.GetThirdByOrderSnReq) (*pb.ThirdResp, error) {
	l := logic.NewGetThirdByOrderSnLogic(ctx, s.svcCtx)
	return l.GetThirdByOrderSn(in)
}

func (s *PayServer) UpdateByOrderSn(ctx context.Context, in *pb.UpdateByOrderSnReq) (*pb.UpdateByOrderSnResp, error) {
	l := logic.NewUpdateByOrderSnLogic(ctx, s.svcCtx)
	return l.UpdateByOrderSn(in)
}
