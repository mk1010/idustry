package main

import (
	"errors"
	"fmt"

	"github.com/mk1010/idustry/common/constant"
	"github.com/mk1010/idustry/config"
	"github.com/mk1010/idustry/handler"
	model "github.com/mk1010/idustry/modules"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var e *gin.Engine

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Sprintf("init config error:%v", err))
	}
	if err := initModel(); err != nil {
		panic(fmt.Sprintf("init model error:%v", err))
	}
	initGin()
	initHandler()
	if err := e.RunTLS(":8080", "./server.crt", "./server.key"); err != nil {
		panic(fmt.Sprintf("gin running error:%v", err))
	}
}

func initGin() {
	if config.ConfInstance.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	e = gin.Default()
	router := gin.Default()
	pprof.Register(router)
	go router.Run(constant.ListenDebugAddr)
}

func initConfig() error {
	curEnv := config.CheckEnv()
	// mk
	configFile := config.GetConfigPath()

	if err := config.Init(configFile); err != nil {
		return err
	}
	if curEnv != config.ConfInstance.Env {
		return errors.New("env error")
	}
	fmt.Printf("Service running in %s mode\n", curEnv)
	return nil
}

func initModel() error {
	return model.Init()
}

func initHandler() {
	handler.Init(e)
}
