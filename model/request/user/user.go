package request

type UserInfoReq struct {
	// 用户名
	Name        	string     	`json:"name" form:"name"`
	// 密码
	Password      	string     	`json:"password" form:"password"`
}

type GetUserListReq struct {
	// 状态
	Status  	int `json:"status" form:"status"`
	// 查询条件
	SearchVal 	string `json:"searchVal" form:"searchVal"`
	// 页码
	Page     	int `json:"page" form:"page"`
	// 每页条数
	PageSize 	int `json:"pageSize" form:"pageSize"`
}

type GetUserInfoReq struct {
	// 用户id
	ID  	int `json:"id" form:"id"`
	Name  	string `json:"name" form:"name"`
}