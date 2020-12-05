package main

import (
	"errors"
	"fmt"
	"industry_identification_center/common/constant"
	"industry_identification_center/config"
	"industry_identification_center/handler"
	"industry_identification_center/model"

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

	initGin()
	if err := initModel(); err != nil {
		panic(fmt.Sprintf("init model error:%v", err))
	}
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
	config.FlagInit()
	configFile := config.Input_ConfDir + "/" + fmt.Sprintf("industry_identification_center_%s.json", curEnv)

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
