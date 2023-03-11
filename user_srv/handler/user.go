package handler

import (
	"context"
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
