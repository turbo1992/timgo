package test

import (
	"fmt"
	"sync"
)

type WebConfig struct {
	port int
}

var cc *WebConfig
var mutex sync.Mutex

func GetConfig() *WebConfig {
	mutex.Lock()
	defer  mutex.Unlock()
	if cc == nil {
		cc = &WebConfig{port: 8080}
	}
	return cc
}

var once sync.Once

func GetSysConfig() *WebConfig {
	once.Do(func() {
		cc = &WebConfig{port: 8080}
	})
	return cc
}

func TestSingleton()  {
	//c1 := GetConfig()
	//c2 := GetConfig()
	//fmt.Println(c1 == c2)

	c1 := GetSysConfig()
	c2 := GetSysConfig()
	fmt.Println(c1 == c2)
}