package initialize

import (
	"tim-go/global"
	"tim-go/model/user"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(
		user.User{},
	)
	global.GVA_LOG.Debug("register table success")
}
