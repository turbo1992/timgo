package global

import (
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"tim-go/config"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG *oplogging.Logger
)

const (
	// 用户状态
	USER_STATE_NORMAL   	= 1
	USER_STATE_FORBIDDEN   	= 2
)
