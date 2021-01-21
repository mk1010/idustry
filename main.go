package main

import (
	"errors"
	"fmt"
	"industry_identification_center/common/constant"
	"industry_identification_center/config"
	"industry_identification_center/handler"
	model "industry_identification_center/modules"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	e *gin.Engine
)

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Sprintf("init config error:%v", err))
	}
	if err := initModel(); err != nil {
		panic(fmt.Sprintf("init model error:%v", err))
	}
	initGin()
	initHandler()
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
