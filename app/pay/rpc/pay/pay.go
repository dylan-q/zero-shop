// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package pay

import (
	"context"

	"zero-shop/app/pay/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetThirdByOrderSnReq = pb.GetThirdByOrderSnReq
	PayReq               = pb.PayReq
	PayResp              = pb.PayResp
	ThirdReq             = pb.ThirdReq
	ThirdResp            = pb.ThirdResp
	UpdateByOrderSnReq   = pb.UpdateByOrderSnReq
	UpdateByOrderSnResp  = pb.UpdateByOrderSnResp

	Pay interface {
		Create(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error)
		Third(ctx context.Context, in *ThirdReq, opts ...grpc.CallOption) (*ThirdResp, error)
		GetThirdByOrderSn(ctx context.Context, in *GetThirdByOrderSnReq, opts ...grpc.CallOption) (*ThirdResp, error)
		UpdateByOrderSn(ctx context.Context, in *UpdateByOrderSnReq, opts ...grpc.CallOption) (*UpdateByOrderSnResp, error)
	}

	defaultPay struct {
		cli zrpc.Client
	}
)

func NewPay(cli zrpc.Client) Pay {
	return &defaultPay{
		cli: cli,
	}
}

func (m *defaultPay) Create(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error) {
	client := pb.NewPayClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultPay) Third(ctx context.Context, in *ThirdReq, opts ...grpc.CallOption) (*ThirdResp, error) {
	client := pb.NewPayClient(m.cli.Conn())
	return client.Third(ctx, in, opts...)
}

func (m *defaultPay) GetThirdByOrderSn(ctx context.Context, in *GetThirdByOrderSnReq, opts ...grpc.CallOption) (*ThirdResp, error) {
	client := pb.NewPayClient(m.cli.Conn())
	return client.GetThirdByOrderSn(ctx, in, opts...)
}

func (m *defaultPay) UpdateByOrderSn(ctx context.Context, in *UpdateByOrderSnReq, opts ...grpc.CallOption) (*UpdateByOrderSnResp, error) {
	client := pb.NewPayClient(m.cli.Conn())
	return client.UpdateByOrderSn(ctx, in, opts...)
}