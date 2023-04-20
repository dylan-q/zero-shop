package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32      `gorm:"primarykey" json:"id"`
	CreatedAt *time.Time `gorm:"column:add_time"`
	UpdateAt  *time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	BirthDay *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"default:male;type:varchar(6) comment 'female 女 male 男'"`
	Role     int        `gorm:"default:1;type:int comment '1 表示普通用户 2 表示管理员'"`
}
