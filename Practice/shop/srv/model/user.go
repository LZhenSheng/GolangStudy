package model

import (
	"time"

	"gorm.io/gorm"
)

// 自定义model
type BaseModel struct {
	ID        int32          `gorm:"primarykey"` //主码
	CreatedAt time.Time      `gorm:"column:add_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt //软删除
	IsDeleted bool
}

/**
1.密文2.密文不可反解
1.对称加密
2.非对称加密
3.md5信息摘要算法
密码如果不可以反解，用户找回密码
*/
// 用户相关表
type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	BirthDay *time.Time `gorm:"type:datetime"` //指针防止保存出错
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment'female表示女,male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户,2表示管理员'"`
}
