package main

import (
	"errors"
	"fmt"
	"industry_identification_center/config"

	"github.com/gin-gonic/gin"
)

var (
	e *gin.Engine
)

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Sprintf("initconfig error:%v", err))
	}

	initGin()

}

func initGin() {
	if config.ConfInstance.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	e = gin.Default()
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
