package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mk1010/idustry/bash"
	"github.com/mk1010/idustry/common/constant"
	"github.com/mk1010/idustry/config"
	"github.com/mk1010/idustry/handler"
	model "github.com/mk1010/idustry/modules"
	"github.com/mk1010/idustry/service"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	"github.com/apache/dubbo-go/common/logger"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	dubboConfig "github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/grpc"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
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
	// initHandler()
	// util.GoSafely(func() {
	// if err := e.RunTLS(":8080", "server.crt", "server.key"); err != nil {
	// panic(fmt.Sprintf("gin running error:%v", err))
	// }
	// }, nil)
	if err := service.Init(context.Background()); err != nil {
		panic(fmt.Sprintf("init nclink service error:%v", err))
	}
	dubboConfig.Load()
	initSignal()
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

func initModel() (err error) {
	err = model.Init()
	return
}

func initHandler() {
	handler.Init(e)
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(3*time.Second), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			logger.Info("provider app exit now...")
			return
		}
	}
}
