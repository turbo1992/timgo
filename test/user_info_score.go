package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type ClassInfo struct {
	Name string
	Count string
}

// 扇入
func fanIn(args ... <-chan interface{}) <-chan interface{} {
	ret := make(chan interface{})
	wg := sync.WaitGroup{}
	for _, ch := range args {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v:= range c {
				ret <- v
			}
		}(ch)
	}
	go func() {
		defer close(ret)
		wg.Wait()
	}()
	return ret
}

func getUserName() <-chan interface{} {
	nameChan := make(chan interface{})
	go func() {
		defer close(nameChan)
		time.Sleep(time.Second*1)
		nameChan <- "<p>name:turbo</p>"
	}()
	return nameChan
}

func getUserScore() <-chan interface{} {
	scoreChan := make(chan interface{})
	go func() {
		defer close(scoreChan)
		time.Sleep(time.Second*3)
		scoreChan <- "<p>score:98</p>"
	}()
	return scoreChan
}

func getUserClass() <-chan interface{} {
	heightChan := make(chan interface{})
	go func() {
		defer close(heightChan)
		time.Sleep(time.Second*5)
		class := ClassInfo{
			Name: "<p>class-01-01</p>",
			Count: "<p>45</p>",
		}
		heightChan <- class
	}()
	return heightChan
}

func TestGetUserInfoWeb() {
	route := gin.New()
	route.GET("/", func(context *gin.Context) {
		// 获取用户姓名、积分、身高
		context.Writer.Header().Add("Transfer-Encoding", "chunked")
		context.Writer.WriteHeader(http.StatusOK)
		ret := fanIn(getUserName(), getUserScore(), getUserClass())
		for v := range ret {
			if getType(v) == "string" {
				context.Writer.Write([]byte(v.(string)))
			}
			if getType(v) == "ClassInfo" {
				//res, _ := json.Marshal(v)
				str := fmt.Sprintf("%v", v)
				context.Writer.Write([]byte(str))
			}
			context.Writer.(http.Flusher).Flush()
		}
	})
	route.Run(":8888")
}

func TestGetUserInfoCode() {
	start := time.Now()
	ret := fanIn(getUserName(), getUserScore(), getUserClass())
	// 扇出
	for v := range ret {
		str := fmt.Sprintf("%v", v)
		fmt.Println(getType(v), str)
	}
	fmt.Printf("用时:%d秒\n", time.Since(start).Milliseconds()/1000)
}

func getType(element interface{}) interface{} {
	switch element.(type) {
	case string:
		return "string"
	case ClassInfo:
		return "ClassInfo"
	}
	return nil
}