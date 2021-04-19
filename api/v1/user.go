package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tim-go/global/response"
	"tim-go/utils"
	userReq "tim-go/model/request/user"
	userSvc "tim-go/service/user"
)

func Register(c *gin.Context)  {
	var info userReq.UserInfoReq
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("注册失败, %v", err), c)
	}

	Verify := utils.Rules{
		"Name": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	VerifyErr := utils.Verify(info, Verify)
	if VerifyErr != nil {
		response.FailWithMessage(VerifyErr.Error(), c)
		return
	}

	err = userSvc.Register(info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("注册失败, %v", err), c)
	} else {
		response.Ok(c)
	}
}

func Login(c *gin.Context)  {
	var info userReq.UserInfoReq
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("登录失败, %v", err), c)
	}

	Verify := utils.Rules{
		"Name": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	VerifyErr := utils.Verify(info, Verify)
	if VerifyErr != nil {
		response.FailWithMessage(VerifyErr.Error(), c)
		return
	}

	err, userInfo := userSvc.Login(info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("登录失败, %v", err), c)
	} else {
		response.OkDetailed(userInfo, "登录成功", c)
	}
}

func UserDetail(c *gin.Context)  {
	var info userReq.GetUserInfoReq
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败, %v", err), c)
	}

	Verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	VerifyErr := utils.Verify(info, Verify)
	if VerifyErr != nil {
		response.FailWithMessage(VerifyErr.Error(), c)
		return
	}

	err, userInfo := userSvc.GetUserDetail(info)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败, %v", err), c)
	} else {
		response.OkWithData(userInfo, c)
	}
}