package rmq

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/mk1010/idustry/config"
)

func TestMain(m *testing.M) {
	curEnv := config.CheckEnv()
	// mk
	configFile := fmt.Sprintf("../../conf/industry_identification_center_%s.json", curEnv)

	if err := config.Init(configFile); err != nil {
		panic(err)
	}
	if curEnv != config.ConfInstance.Env {
		panic(errors.New("env error"))
	}
	fmt.Printf("Service running in %s mode\n", curEnv)
	m.Run()
}

func TestRmq(t *testing.T) {
	t.Log(Init())
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver(config.ConfInstance.RMQNamingService)),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	err = c.Start()
	if err != nil {
		t.Fatal(err)
	}
	err = c.Shutdown()
	if err != nil {
		t.Fatal(err)
	}
}
