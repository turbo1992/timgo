package test

import "fmt"

type logClosure func(format string, v...interface{})

func LoggerWrapper(logType string) logClosure {
	fmt.Println("1111111")
	return func(format string, v...interface{}) {
		fmt.Printf(fmt.Sprintf("[%s] %s", logType, format), v...,)
		fmt.Println("")
		fmt.Println("2222222")
	}
}

func TestCallBack()  {
	info_logger := LoggerWrapper("INFO")
	info_logger("this is a %s log", "info")
}