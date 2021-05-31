package bash

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

const (
	outputDir   = "./logs/"
	logFileName = "agent.log"
)

func init() {
	err := os.Setenv("prod", "true")
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(outputDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(outputDir, os.ModePerm)
			if err != nil {
				panic(fmt.Sprintf("mkdir failed![%v]\n", err))
			}
		}
	}
	zapLogConf := new(zap.Config)
	err = yaml.Unmarshal([]byte(logConf), zapLogConf)
	if err != nil {
		panic(err)
	}
	logHook := getWriter(logFileName)
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(zapLogConf.EncoderConfig), zapcore.AddSync(logHook), zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return true
		})),
	)
	zapLogger := zap.New(core, zap.AddCallerSkip(1))
	logger.SetLogger(zapLogger.Sugar())
	logConf = "" // 释放内存,可能也许

	dubboProviderConf := new(config.ProviderConfig)
	f, err := ioutil.ReadFile("./conf/server.yml")
	if err != nil {
		panic("未找到dubbo服务配置文件")
	}
	err = yaml.Unmarshal(f, dubboProviderConf)
	if err != nil {
		panic(err)
	}
	config.SetProviderConfig(*dubboProviderConf)
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

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		// 没有使用go风格反人类的format格式
		outputDir+filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
