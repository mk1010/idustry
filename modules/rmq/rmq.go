package rmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/mk1010/idustry/config"
)

var RmqProduce rocketmq.Producer

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
