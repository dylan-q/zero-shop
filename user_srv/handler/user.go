package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/model"
	"mxshop_srvs/user_srv/proto"
)

type UserServer struct{}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfo := proto.UserInfoResponse{
		Id:       uint32(user.ID),
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.BirthDay != nil {
		userInfo.BirthDay = uint64(user.BirthDay.Unix())
	}
	return userInfo
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func (u *UserServer) GetUser(ctx context.Context, info *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User

	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error

	}
	rsp := &proto.UserListResponse{}

	rsp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(info.Pn), int(info.PSize))).Find(&users)
	for _, v := range users {
		userInfoRsp := ModelToResponse(v)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

func (u *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfo := ModelToResponse(user)
	return &userInfo, nil
}

func (u *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfo := ModelToResponse(user)
	return &userInfo, nil
}

func (u *UserServer) CreateUser(ctx context.Context, info *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: info.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Mobile = info.Mobile
	user.NickName = info.NickName
}
