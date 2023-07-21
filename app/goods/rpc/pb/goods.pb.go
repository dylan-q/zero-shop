// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: app/goods/rpc/goods.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_goods_rpc_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_goods_rpc_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_app_goods_rpc_goods_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type CategoryInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CategoryInfoResponse) Reset() {
	*x = CategoryInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_goods_rpc_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryInfoResponse) ProtoMessage() {}

func (x *CategoryInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_goods_rpc_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryInfoResponse.ProtoReflect.Descriptor instead.
func (*CategoryInfoResponse) Descriptor() ([]byte, []int) {
	return file_app_goods_rpc_goods_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*GoodsInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_goods_rpc_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_goods_rpc_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_app_goods_rpc_goods_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetData() []*GoodsInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type GoodsInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId  int64                 `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	GoodsName   string                `protobuf:"bytes,3,opt,name=goodsName,proto3" json:"goodsName,omitempty"`
	GoodsDetail string                `protobuf:"bytes,4,opt,name=goodsDetail,proto3" json:"goodsDetail,omitempty"`
	ClickNum    int64                 `protobuf:"varint,5,opt,name=clickNum,proto3" json:"clickNum,omitempty"`
	SoldNum     int64                 `protobuf:"varint,6,opt,name=soldNum,proto3" json:"soldNum,omitempty"`
	FavNum      int64                 `protobuf:"varint,7,opt,name=favNum,proto3" json:"favNum,omitempty"`
	MarketPrice float32               `protobuf:"fixed32,8,opt,name=marketPrice,proto3" json:"marketPrice,omitempty"`
	ShopPrice   float32               `protobuf:"fixed32,9,opt,name=shopPrice,proto3" json:"shopPrice,omitempty"`
	GoodsBrief  string                `protobuf:"bytes,10,opt,name=goodsBrief,proto3" json:"goodsBrief,omitempty"`
	ShipFree    int64                 `protobuf:"varint,11,opt,name=shipFree,proto3" json:"shipFree,omitempty"`
	IsNew       int64                 `protobuf:"varint,12,opt,name=isNew,proto3" json:"isNew,omitempty"`
	IsHot       int64                 `protobuf:"varint,13,opt,name=isHot,proto3" json:"isHot,omitempty"`
	OnSale      int64                 `protobuf:"varint,14,opt,name=onSale,proto3" json:"onSale,omitempty"`
	CreateTime  int64                 `protobuf:"varint,15,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Category    *CategoryInfoResponse `protobuf:"bytes,16,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GoodsInfoResponse) Reset() {
	*x = GoodsInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_goods_rpc_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfoResponse) ProtoMessage() {}

func (x *GoodsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_goods_rpc_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfoResponse.ProtoReflect.Descriptor instead.
func (*GoodsInfoResponse) Descriptor() ([]byte, []int) {
	return file_app_goods_rpc_goods_proto_rawDescGZIP(), []int{3}
}

func (x *GoodsInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsInfoResponse) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GoodsInfoResponse) GetGoodsName() string {
	if x != nil {
		return x.GoodsName
	}
	return ""
}

func (x *GoodsInfoResponse) GetGoodsDetail() string {
	if x != nil {
		return x.GoodsDetail
	}
	return ""
}

func (x *GoodsInfoResponse) GetClickNum() int64 {
	if x != nil {
		return x.ClickNum
	}
	return 0
}

func (x *GoodsInfoResponse) GetSoldNum() int64 {
	if x != nil {
		return x.SoldNum
	}
	return 0
}

func (x *GoodsInfoResponse) GetFavNum() int64 {
	if x != nil {
		return x.FavNum
	}
	return 0
}

func (x *GoodsInfoResponse) GetMarketPrice() float32 {
	if x != nil {
		return x.MarketPrice
	}
	return 0
}

func (x *GoodsInfoResponse) GetShopPrice() float32 {
	if x != nil {
		return x.ShopPrice
	}
	return 0
}

func (x *GoodsInfoResponse) GetGoodsBrief() string {
	if x != nil {
		return x.GoodsBrief
	}
	return ""
}

func (x *GoodsInfoResponse) GetShipFree() int64 {
	if x != nil {
		return x.ShipFree
	}
	return 0
}

func (x *GoodsInfoResponse) GetIsNew() int64 {
	if x != nil {
		return x.IsNew
	}
	return 0
}

func (x *GoodsInfoResponse) GetIsHot() int64 {
	if x != nil {
		return x.IsHot
	}
	return 0
}

func (x *GoodsInfoResponse) GetOnSale() int64 {
	if x != nil {
		return x.OnSale
	}
	return 0
}

func (x *GoodsInfoResponse) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GoodsInfoResponse) GetCategory() *CategoryInfoResponse {
	if x != nil {
		return x.Category
	}
	return nil
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_goods_rpc_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_goods_rpc_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_app_goods_rpc_goods_proto_rawDescGZIP(), []int{4}
}

func (x *DetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_app_goods_rpc_goods_proto protoreflect.FileDescriptor

var file_app_goods_rpc_goods_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x3d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x3a,
	0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe7, 0x03, 0x0a, 0x11,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x6f, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x6f, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x76, 0x4e, 0x75,
	0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x61, 0x76, 0x4e, 0x75, 0x6d, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x72, 0x69, 0x65, 0x66, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x72, 0x69, 0x65, 0x66, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x46, 0x72, 0x65, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x46, 0x72, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x73, 0x4e, 0x65, 0x77, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x73, 0x4e, 0x65,
	0x77, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x53, 0x61, 0x6c,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x6e, 0x53, 0x61, 0x6c, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x66, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12,
	0x29, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_goods_rpc_goods_proto_rawDescOnce sync.Once
	file_app_goods_rpc_goods_proto_rawDescData = file_app_goods_rpc_goods_proto_rawDesc
)

func file_app_goods_rpc_goods_proto_rawDescGZIP() []byte {
	file_app_goods_rpc_goods_proto_rawDescOnce.Do(func() {
		file_app_goods_rpc_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_goods_rpc_goods_proto_rawDescData)
	})
	return file_app_goods_rpc_goods_proto_rawDescData
}

var file_app_goods_rpc_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_goods_rpc_goods_proto_goTypes = []interface{}{
	(*ListRequest)(nil),          // 0: pb.ListRequest
	(*CategoryInfoResponse)(nil), // 1: pb.CategoryInfoResponse
	(*ListResponse)(nil),         // 2: pb.ListResponse
	(*GoodsInfoResponse)(nil),    // 3: pb.GoodsInfoResponse
	(*DetailRequest)(nil),        // 4: pb.DetailRequest
}
var file_app_goods_rpc_goods_proto_depIdxs = []int32{
	3, // 0: pb.ListResponse.data:type_name -> pb.GoodsInfoResponse
	1, // 1: pb.GoodsInfoResponse.category:type_name -> pb.CategoryInfoResponse
	0, // 2: pb.goods.List:input_type -> pb.ListRequest
	4, // 3: pb.goods.Detail:input_type -> pb.DetailRequest
	2, // 4: pb.goods.List:output_type -> pb.ListResponse
	3, // 5: pb.goods.Detail:output_type -> pb.GoodsInfoResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_goods_rpc_goods_proto_init() }
func file_app_goods_rpc_goods_proto_init() {
	if File_app_goods_rpc_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_goods_rpc_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_goods_rpc_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_goods_rpc_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_goods_rpc_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_goods_rpc_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_goods_rpc_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_goods_rpc_goods_proto_goTypes,
		DependencyIndexes: file_app_goods_rpc_goods_proto_depIdxs,
		MessageInfos:      file_app_goods_rpc_goods_proto_msgTypes,
	}.Build()
	File_app_goods_rpc_goods_proto = out.File
	file_app_goods_rpc_goods_proto_rawDesc = nil
	file_app_goods_rpc_goods_proto_goTypes = nil
	file_app_goods_rpc_goods_proto_depIdxs = nil
}
