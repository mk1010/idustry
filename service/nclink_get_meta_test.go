package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/mk1010/idustry/config"
	"github.com/mk1010/idustry/modules/industry_identification_center/model"
	"github.com/mk1010/industry_adaptor/nclink"
	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	curEnv := config.CheckEnv()
	// mk
	configFile := fmt.Sprintf("../conf/industry_identification_center_%s.json", curEnv)

	if err := config.Init(configFile); err != nil {
		panic(err)
	}
	if err := model.Init(); err != nil {
		panic(err)
	}
	if curEnv != config.ConfInstance.Env {
		panic(errors.New("env error"))
	}
	fmt.Printf("Service running in %s mode\n", curEnv)
	m.Run()
}

func TestGetMeta(t *testing.T) {
	b := NewNcLinkServiceProvider()
	resp, err := b.NCLinkGetMeta(context.Background(), &nclink.NCLinkMetaDataReq{
		AdaptorId: []string{"ada_mock_23"},
	})
	t.Log(resp, err)
}
