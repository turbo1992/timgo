package initialize

import (
	"fmt"
	"tim-go/global"
	"tim-go/model/user"

)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(
		user.User{},
	)
	fmt.Println("2222222222")
	global.GVA_LOG.Debug("register table success")
	fmt.Println("33333333333")

}
