package task

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mk1010/idustry/config"
	"github.com/mk1010/industry_adaptor/nclink"
)

func TestMain(m *testing.M) {
	curEnv := config.CheckEnv()
	// mk
	configFile := fmt.Sprintf("../conf/industry_identification_center_%s.json", curEnv)

	if err := config.Init(configFile); err != nil {
		panic(err)
	}
	if curEnv != config.ConfInstance.Env {
		panic(errors.New("env error"))
	}
	fmt.Printf("Service running in %s mode\n", curEnv)
	m.Run()
}

func TestNCLinkCommandTopic(t *testing.T) {
	t.Log(NCLinkCommandTopic(&nclink.NCLinkTopicSub{
		Topic: nclink.CommandTopic,
	}, nil))
}
