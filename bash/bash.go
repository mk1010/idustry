package bash

import (
	"net"
	"os"
	"strings"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/config"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func init() {
	err := os.Setenv("prod", "true")
	if err != nil {
		panic(err)
	}

	zapLogConf := new(zap.Config)
	err = yaml.Unmarshal([]byte(logConf), zapLogConf)
	if err != nil {
		panic(err)
	}
	logger.InitLogger(zapLogConf)
	logConf = "" // 释放内存
	dubboProviderConf := new(config.ProviderConfig)
	err = yaml.Unmarshal([]byte(providerConfig), dubboProviderConf)
	if err != nil {
		panic(err)
	}

	config.SetProviderConfig(*dubboProviderConf)
	providerConfig = ""
}

func getPublicIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		logger.Error(err)
		return ""
	}
	defer conn.Close()
	s := conn.LocalAddr().String()
	if index := strings.LastIndex(s, ":"); index != -1 {
		s = s[:index]
	}
	return s
}