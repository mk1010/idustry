package handler

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(e *gin.Engine)
}

var (
	handlers []Handler
)

func Init(e *gin.Engine) {
	handlers = make([]Handler, 0)
	handlers = append(handlers, &homePage{})
	registerHandlers(e)
}

func registerHandlers(e *gin.Engine) {
	for _, h := range handlers {
		h.Register(e)
	}

	routes := e.Routes()
	for _, route := range routes {
		registerHystrix(route)
	}
}

func registerHystrix(route gin.RouteInfo) {
	cmdConf := hystrix.CommandConfig{
		Timeout:                3000,
		MaxConcurrentRequests:  5000,
		RequestVolumeThreshold: 100,
		SleepWindow:            30000,
		ErrorPercentThreshold:  50,
	}
	//
	//todo 查看断路器工作方式，各path定制化断路器配置
	hystrix.ConfigureCommand(route.Path, cmdConf)
}
