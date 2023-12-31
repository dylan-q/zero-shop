// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package order

import (
	"context"

	"zero-shop/app/order/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateOrderReq        = pb.CreateOrderReq
	CreateOrderResp       = pb.CreateOrderResp
	DetailReq             = pb.DetailReq
	GoodsInfo             = pb.GoodsInfo
	ListReq               = pb.ListReq
	ListResp              = pb.ListResp
	OrderInfoResp         = pb.OrderInfoResp
	ThirdInfo             = pb.ThirdInfo
	UpdateOrderStatusReq  = pb.UpdateOrderStatusReq
	UpdateOrderStatusResp = pb.UpdateOrderStatusResp

	Order interface {
		Create(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
		List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
		Detail(ctx context.Context, in *DetailReq, opts ...grpc.CallOption) (*OrderInfoResp, error)
		UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Create(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultOrder) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultOrder) Detail(ctx context.Context, in *DetailReq, opts ...grpc.CallOption) (*OrderInfoResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultOrder) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.UpdateOrderStatus(ctx, in, opts...)
}
