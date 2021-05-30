package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/mk1010/idustry/config"
	"github.com/mk1010/idustry/modules/industry_identification_center/model"
	nclink_data "github.com/mk1010/idustry/modules/industry_identification_center/model/property/nclinkdevicedata"
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
		ComponentId: []string{"component_mock_1"},
	})
	t.Log(resp, err)
}

func TestSendData(t *testing.T) {
	b := NewNcLinkServiceProvider()
	resp, err := b.NCLinkSendData(context.Background(), &nclink.NCLinkDataMessage{
		DataId:      "data_id_test_1",
		AdaptorId:   "ada_mock_23",
		DeviceId:    "device_id_test_1",
		ComponentId: "component_id_test_1",
		DataItemId:  "data_item_id_test_1",
		Payloads: []*nclink.NCLinkPayloads{
			{
				UnixTimeMs: 0,
				Payload:    []byte("hello mk"),
			},
		},
	})
	t.Log(resp, err)
}

func TestCollectData(tt *testing.T) {
	ans := make([]string, 0)
	for t := uint64(1622027556); t <= uint64(1622027584); t++ {
		dataModel := make([]*model.NclinkDeviceData, 0)
		err := model.DeviceInfoDB.WriteDB().Table(nclink_data.Namespace).Select(nclink_data.AllFileds).Where(nclink_data.AdaptorTime+" >= ? and "+nclink_data.AdaptorTime+" <= ?",
			t*1000, t*1000+999).Find(&dataModel).Error
		if err != nil {
			tt.Fatal(err)
		}
		ans = append(ans, fmt.Sprintf("%v,%v", t, len(dataModel)))
	}
	tt.Log(ans)
}

func TestSlice(t *testing.T) {
	m := make([]int, 0)
	for i := 0; i < 10; i++ {
		func(s []int) {
			// *s = append(*s, i)
			s = append(s, i)
		}(m)
	}
	t.Log(m)
}
