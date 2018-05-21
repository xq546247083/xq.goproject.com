package main

import (
	"fmt"
)

func main() {
	// 新建对象
	errorLoggerObj := &errorLogger{}
	debugLoggerObj := &debugLogger{}
	infoLoggerObj := &infoLogger{}
	baseLogger := &logger{}

	// 错误日志器没有重写Log方法，导致调用的基类的Log
	errorLoggerObj.Log("A")

	//
	errorLoggerObj.setNextLogger(debugLoggerObj)
	errorLoggerObj.Log("B")

	debugLoggerObj.setNextLogger(infoLoggerObj)
	errorLoggerObj.Log("C")

	infoLoggerObj.setNextLogger(baseLogger)
	errorLoggerObj.Log("D")
}

// ----------------日志基类----------------

// 日志接口
type iLogger interface {
	setNextLogger(iLogger)
	Log(string)
}

//  日志基类
type logger struct {
	nextLogger iLogger
}

// 设置下一个级别的日志记录这
func (this *logger) setNextLogger(logger iLogger) {
	this.nextLogger = logger
}

func (this *logger) Log(info string) {
	if this.nextLogger != nil {
		this.nextLogger.Log(info)
		return
	}

	fmt.Println("baseLog", info)
}

// ------------各种各样级别的日志-----------------

//  错误日志
type errorLogger struct {
	logger
}

//  调试日志
type debugLogger struct {
	logger
}

// 经过debug责任链的日志，都要记录一下
func (this *debugLogger) Log(info string) {
	fmt.Println("debugLog", info)

	if this.nextLogger != nil {
		this.nextLogger.Log(info)
	}
}

//  消息日志
type infoLogger struct {
	logger
}

func (this *infoLogger) Log(info string) {
	if this.nextLogger != nil {
		this.nextLogger.Log(info)
		return
	}

	fmt.Println("infoLog", info)
}
