package core

import (
	"flag"
	"fmt"
	"tim-go/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	defaultConfigFile := flag.String("f", "config.yaml", "number")
	flag.Parse()
	v := viper.New()
	v.SetConfigFile(*defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GVA_VP = v
}
