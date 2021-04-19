package user

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name  string `json:"name" form:"name" gorm:"column:name;comment:'用户名';type:varchar(100)"`
	Password  string `json:"password" form:"password" gorm:"column:password;comment:'密码';type:varchar(100)"`
	State  int `json:"state" form:"state" gorm:"column:state;comment:'用户状态';type:int(10)"`
}

func (User) TableName() string {
	return "user"
}