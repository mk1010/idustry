package common

import (
	"strconv"
	"strings"

	"github.com/mk1010/idustry/common/constant"
	"github.com/mk1010/idustry/modules/industry_identification_center/model"
	"github.com/mk1010/industry_adaptor/nclink"
)

func AdaptorMetaModelToProto(m *model.AdaptorMeta) *nclink.NCLinkAdaptor {
	deviceID := strings.Split(m.DeviceID, constant.SplitChar)
	return &nclink.NCLinkAdaptor{
		AdaptorId:   m.AdaptorID,
		Name:        m.Name,
		AdaptorType: m.AdaptorType,
		Description: m.Description,
		DeviceId:    deviceID,
		Config:      []byte(m.Config),
	}
}

func DevicerMetaModelToProto(m *model.DeviceMeta) *nclink.NCLinkDevice {
	componentId := strings.Split(m.ComponentID, constant.SplitChar)
	return &nclink.NCLinkDevice{
		DeviceId:    m.DeviceID,
		Name:        m.Name,
		DeviceType:  m.DeviceType,
		Description: m.Description,
		DeviceGroup: m.DeviceGroup,
		ComponentId: componentId,
		Config:      []byte(m.Config),
	}
}

func ComponentMetaModelToProto(m *model.ComponentMeta) *nclink.NCLinkComponent {
	return &nclink.NCLinkComponent{
		ComponentId:   m.ComponentID,
		Name:          m.Name,
		ComponentType: m.ComponentType,
		Description:   m.Description,
		Config:        []byte(m.Config),
	}
}

func DataItemMetaModelToProto(m *model.DataItemMeta) *nclink.NCLinkDataItem {
	items := strings.Split(m.Items, constant.SplitChar)
	dataUnit := strings.Split(m.DataUnit, constant.SplitChar)
	if len(items)&1 != 0 || len(dataUnit)&1 != 0 {
		return nil
	}
	nclinkDataItemMin := make([]*nclink.NCLinkDataItemMin, 0, len(items)/2)
	for i := 0; i < len(items); i += 2 {
		kind, err := strconv.ParseInt(items[i+1], 10, 32)
		if err != nil {
			return nil
		}
		m := &nclink.NCLinkDataItemMin{
			FiledName: items[i],
			Kind:      nclink.DataKind(kind),
		}
		nclinkDataItemMin = append(nclinkDataItemMin, m)
	}
	dataUnitMap := make(map[string]string, len(dataUnit)/2)
	for i := 0; i < len(dataUnit); i += 2 {
		dataUnitMap[dataUnit[i]] = dataUnit[i+1]
	}
	return &nclink.NCLinkDataItem{
		DataItemId:   m.DataItemID,
		Name:         m.Name,
		DataItemType: m.DataItemType,
		Description:  m.Description,
		Items:        nclinkDataItemMin,
		DataUnit:     dataUnitMap,
	}
}

func SamplingInfoMetaModelToProto(m *model.SamplingInfoMeta) *nclink.NCLinkSampleInfo {
	return &nclink.NCLinkSampleInfo{
		SampleInfoId:   m.SamplingInfoID,
		SampleInfoType: m.SamplingInfoType,
		SamplingPeriod: m.SamplingPeriod,
		UploadPeriod:   m.UploadPeriod,
	}
}
