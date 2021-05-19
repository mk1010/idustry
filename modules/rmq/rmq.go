package rmq

import (
	"fmt"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/mk1010/idustry/config"
)

var RmqProduce rocketmq.Producer

func init() {
	rlog.SetLogger(MockLogger{})
}

func Init() error {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(config.ConfInstance.RMQNamingService),
		producer.WithRetry(2))
	if err != nil {
		return err
	}
	err = p.Start()
	if err != nil {
		return err
	}
	RmqProduce = p

	return nil
}

type MockLogger struct{}

func (m MockLogger) Debug(msg string, fields map[string]interface{}) {
	logger.Debug(msg, fields)
}

func (m MockLogger) Info(msg string, fields map[string]interface{}) {
	logger.Info(msg, fields)
}

func (m MockLogger) Warning(msg string, fields map[string]interface{}) {
	logger.Warn(msg, fields)
}

func (m MockLogger) Error(msg string, fields map[string]interface{}) {
	logger.Error(msg, fields)
}

func (m MockLogger) Fatal(msg string, fields map[string]interface{}) {
	logger.Error(msg, fields)
	panic(fmt.Errorf("%s %v", msg, fields))
}

func (m MockLogger) Level(level string) {
}

func (m MockLogger) OutputPath(path string) (err error) {
	return
}
