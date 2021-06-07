package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/mk1010/idustry/common/constant"
	"github.com/mk1010/idustry/modules/industry_identification_center/model"
	adaptor_meta "github.com/mk1010/idustry/modules/industry_identification_center/model/property/adaptormeta"
	component_meta "github.com/mk1010/idustry/modules/industry_identification_center/model/property/componentmeta"
	dataitem_meta "github.com/mk1010/idustry/modules/industry_identification_center/model/property/dataitemmeta"
	device_meta "github.com/mk1010/idustry/modules/industry_identification_center/model/property/devicemeta"
	samplinginfo_meta "github.com/mk1010/idustry/modules/industry_identification_center/model/property/samplinginfometa"
	"github.com/mk1010/idustry/service/common"
	"github.com/mk1010/industry_adaptor/nclink"
)

func (n *NcLinkServiceProvider) NCLinkGetMeta(ctx context.Context, req *nclink.NCLinkMetaDataReq) (*nclink.NCLinkMetaDataResp, error) {
	logger.Info("NCLinkGetMeta接口收到请求:", req)
	// todo  参数检查 防止被恶意调用
	resp, err := n.nclinkGetMetaImpl(ctx, req)
	if err != nil {
		resp = new(nclink.NCLinkMetaDataResp)
		resp.BaseResp = &nclink.NCLinkBaseResp{
			StatusCode: nclink.DBErrorCommon,
			Detail:     err.Error(),
		}
		return resp, nil
	}
	resp.BaseResp = &nclink.NCLinkBaseResp{
		StatusCode: nclink.StatusOk,
	}
	return resp, nil
}

func (n *NcLinkServiceProvider) nclinkGetMetaImpl(ctx context.Context, req *nclink.NCLinkMetaDataReq) (*nclink.NCLinkMetaDataResp, error) {
	resp := new(nclink.NCLinkMetaDataResp)
	if len(req.AdaptorId) > 0 {
		adaptorModelMeta := make([]*model.AdaptorMeta, 0, len(req.AdaptorId))
		err := model.DeviceInfoDB.ReadDB().Table(adaptor_meta.Namespace).Select(adaptor_meta.AllFileds).Where(fmt.Sprintf("%s in ? and %s = ?",
			adaptor_meta.AdaptorID, adaptor_meta.DeleteTime), req.AdaptorId, constant.NotDeleteTime).Find(&adaptorModelMeta).Error
		if err != nil {
			logger.Error("数据库查询错误 适配器元数据 err:", err)
			return nil, err
		}
		adaptorMeta := make([]*nclink.NCLinkAdaptor, 0, len(req.AdaptorId))
		for _, modelMeta := range adaptorModelMeta {
			nclinkMeta := common.AdaptorMetaModelToProto(modelMeta)
			if nclinkMeta == nil {
				continue
			}
			adaptorMeta = append(adaptorMeta, nclinkMeta)
		}
		resp.Adaptors = adaptorMeta
	}
	if len(req.DeviceId) > 0 {
		deviceModelMeta := make([]*model.DeviceMeta, 0, len(req.DeviceId))
		err := model.DeviceInfoDB.ReadDB().Table(device_meta.Namespace).Select(device_meta.AllFileds).Where(fmt.Sprintf("%s in ? and %s = ?", device_meta.DeviceID,
			device_meta.DeleteTime), req.DeviceId, constant.NotDeleteTime).Find(&deviceModelMeta).Error
		if err != nil {
			logger.Error("数据库查询错误 设备元数据 err:", err)
			return nil, err
		}
		deviceMeta := make([]*nclink.NCLinkDevice, 0, len(req.DeviceId))
		for _, modelMeta := range deviceModelMeta {
			nclinkMeta := common.DevicerMetaModelToProto(modelMeta)
			if nclinkMeta == nil {
				continue
			}
			deviceMeta = append(deviceMeta, nclinkMeta)
		}
		resp.Devices = deviceMeta
	}
	if len(req.ComponentId) > 0 {
		componentModelMeta := make([]*model.ComponentMeta, 0, len(req.ComponentId))
		err := model.DeviceInfoDB.ReadDB().Table(component_meta.Namespace).Select(component_meta.AllFileds).Where(fmt.Sprintf("%s in ? and %s = ?", component_meta.ComponentID,
			component_meta.DeleteTime), req.ComponentId, constant.NotDeleteTime).Find(&componentModelMeta).Error
		if err != nil {
			logger.Error("数据库查询错误 组件元数据 err:", err)
			return nil, err
		}
		componentMeta := make([]*nclink.NCLinkComponent, 0, len(req.ComponentId))
		for _, modelMeta := range componentModelMeta {
			nclinkMeta := common.ComponentMetaModelToProto(modelMeta)
			if nclinkMeta == nil {
				continue
			}
			dataItemIDs := strings.Split(modelMeta.DataItemID, constant.SplitChar)
			samplingInfoIDs := strings.Split(modelMeta.SampleInfoID, constant.SplitChar)
			if len(dataItemIDs) != len(samplingInfoIDs) {
				logger.Errorf("组件ID:%s 元数据错误，数据项不匹配", modelMeta.ComponentID)
			}
			res, err := n.NCLinkGetMeta(ctx, &nclink.NCLinkMetaDataReq{
				DataItemId:   dataItemIDs,
				SampleInfoId: samplingInfoIDs,
			})
			if err != nil {
				return nil, err
			}
			dataItemMap := make(map[string]*nclink.NCLinkDataItem, len(dataItemIDs))
			for _, dataItem := range res.DataItems {
				dataItemMap[dataItem.DataItemId] = dataItem
			}
			samplingInfoMap := make(map[string]*nclink.NCLinkSampleInfo, len(samplingInfoIDs))
			for _, samplingInfo := range res.SampleInfos {
				samplingInfoMap[samplingInfo.SampleInfoId] = samplingInfo
			}
			dataInfo := make([]*nclink.NCLinkDataInfo, 0, len(dataItemIDs))
			for i := range dataItemIDs {
				dataItem, ok := dataItemMap[dataItemIDs[i]]
				if !ok {
					continue
				}
				samplingInfo, ok := samplingInfoMap[samplingInfoIDs[i]]
				if !ok {
					continue
				}
				dataInfo = append(dataInfo, &nclink.NCLinkDataInfo{
					DataItem:   dataItem,
					SampleInfo: samplingInfo,
				})
			}
			nclinkMeta.DataInfo = dataInfo
			componentMeta = append(componentMeta, nclinkMeta)
		}
		resp.Components = componentMeta
	}
	if len(req.DataItemId) > 0 {
		dataItemModelMeta := make([]*model.DataItemMeta, 0, len(req.DataItemId))
		err := model.DeviceInfoDB.ReadDB().Table(dataitem_meta.Namespace).Select(dataitem_meta.AllFileds).Where(fmt.Sprintf("%s in ? and %s = ?", dataitem_meta.DataItemID,
			dataitem_meta.DeleteTime), req.DataItemId, constant.NotDeleteTime).Find(&dataItemModelMeta).Error
		if err != nil {
			logger.Error("数据库查询错误 数据项元数据 err:", err)
			return nil, err
		}
		dataItemMeta := make([]*nclink.NCLinkDataItem, 0, len(req.DataItemId))
		for _, modelMeta := range dataItemModelMeta {
			nclinkMeta := common.DataItemMetaModelToProto(modelMeta)
			if nclinkMeta == nil {
				continue
			}

			dataItemMeta = append(dataItemMeta, nclinkMeta)
		}
		resp.DataItems = dataItemMeta
	}
	if len(req.SampleInfoId) > 0 {
		samplingInfoModelMeta := make([]*model.SamplingInfoMeta, 0, len(req.SampleInfoId))
		err := model.DeviceInfoDB.ReadDB().Table(samplinginfo_meta.Namespace).Select(samplinginfo_meta.AllFileds).Where(fmt.Sprintf("%s in ? and %s = ?", samplinginfo_meta.SamplingInfoID,
			samplinginfo_meta.DeleteTime), req.SampleInfoId, constant.NotDeleteTime).Find(&samplingInfoModelMeta).Error
		if err != nil {
			logger.Error("数据库查询错误 采样元数据 err:", err)
			return nil, err
		}
		samplingInfoMeta := make([]*nclink.NCLinkSampleInfo, 0, len(req.SampleInfoId))
		for _, modelMeta := range samplingInfoModelMeta {
			nclinkMeta := common.SamplingInfoMetaModelToProto(modelMeta)
			if nclinkMeta == nil {
				continue
			}

			samplingInfoMeta = append(samplingInfoMeta, nclinkMeta)
		}
		resp.SampleInfos = samplingInfoMeta
	}
	logger.Info("NCLinkGetMeta接口收到请求:", req, "响应:", resp)
	return resp, nil
}
