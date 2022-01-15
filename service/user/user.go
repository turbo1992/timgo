package user

import (
	"errors"
	"fmt"
	"sync"
	"tim-go/global"
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

func AddUser() error {
	wg := sync.WaitGroup{}
	for i:=0; i<1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			userReturn := userModel.User{
				Name: "user" + string(i),
				Password: "userpwd",
				State: global.USER_STATE_NORMAL,
			}
			global.GVA_DB.Create(&userReturn)
		}()
	}
	wg.Wait()
	return nil
}

func Find() (error, interface{})  {
	wg := sync.WaitGroup{}
	idList := []string{}
	for i:=0; i<100; i++ {
		num := 100 + i
		wg.Add(1)
		go func() {
			defer wg.Done()
			userInfo := userModel.User{}
			global.GVA_DB.Where("id = ?", num).First(&userInfo)
			idList = append(idList, fmt.Sprintf("%s", userInfo.Name))
		}()
	}
	wg.Wait()
	fmt.Println("个数", len(idList))
	return nil, idList
}