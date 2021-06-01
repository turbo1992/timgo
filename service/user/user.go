package user

import (
	"tim-go/global"
	"errors"
	userReq "tim-go/model/request/user"
	userModel "tim-go/model/user"
)

func Register(req userReq.UserInfoReq) error {
	var user userModel.User
	notRegister := global.GVA_DB.Where("name = ?", req.Name).First(&user).RecordNotFound()
	if !notRegister {
		return errors.New("用户名已注册")
	}

	userReturn := userModel.User{
		Name: req.Name,
		Password: req.Password,
		State: global.USER_STATE_NORMAL,
	}
	err := global.GVA_DB.Create(&userReturn).Error
	if err != nil {
		return errors.New("注册失败")
	}
	return err
}

func Login(req userReq.UserInfoReq) (error, userModel.User) {
	userInfo := userModel.User{}
	err := global.GVA_DB.Where("name = ? AND password = ?", req.Name, req.Password).First(&userInfo).Error
	if err != nil {
		return errors.New("用户名或密码有误"), userInfo
	}
	if userInfo.State == global.USER_STATE_FORBIDDEN {
		userInfo = userModel.User{}
		return errors.New("用户已被禁用，请联系管理员"), userInfo
	}
	return err, userInfo
}

func GetUserDetail(req userReq.GetUserInfoReq) (error, userModel.User)  {
	userInfo := userModel.User{}
	global.GVA_DB.Row()
	err := global.GVA_DB.Where("id = ?", req.ID).First(&userInfo).Error
	if err != nil {
		return errors.New("查询失败"), userInfo
	}
	return err, userInfo
}