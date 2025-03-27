package module

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid        uint64         `gorm:"column:uid;primaryKey;auto_increment" json:"uid"`
	UserName   string         `gorm:"column:username;not null" json:"user_name"`
	Passwd     string         `gorm:"column:passwd;not null" json:"passwd"`
	Email      string         `gorm:"column:email;not null" json:"email"`
	Icon       string         `gorm:"column:icon;not null" json:"icon"`
	NickName   string         `gorm:"column:nickname;not null" json:"nick_name"`
	Salt       string         `gorm:"column:salt;not null" json:"salt"`
	CreatedAT  time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	Updated_at time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAT  gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type Tabler interface {
	TableName() string
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetCacheKey() string {
	return fmt.Sprintf("%s:%d", u.TableName(), u.Uid)
}
