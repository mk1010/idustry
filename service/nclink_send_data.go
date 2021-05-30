package service

import (
	"context"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/mk1010/idustry/modules/industry_identification_center/model"
	nclink_data "github.com/mk1010/idustry/modules/industry_identification_center/model/property/nclinkdevicedata"
	"github.com/mk1010/industry_adaptor/nclink"
)

func (n *NcLinkServiceProvider) NCLinkSendData(ctx context.Context, req *nclink.NCLinkDataMessage) (*nclink.NCLinkBaseResp, error) {
	logger.Info("NCLinkSendData接口收到请求:", req)
	modelData := make([]*model.NclinkDeviceData, 0, len(req.Payloads))
	for _, msg := range req.Payloads {
		data := &model.NclinkDeviceData{
			DataID:      req.DataId,
			AdaptorID:   req.AdaptorId,
			DeviceID:    req.DeviceId,
			ComponentID: req.ComponentId,
			DataItemID:  req.DataItemId,
			Payload:     string(msg.Payload),
			AdaptorTime: msg.UnixTimeMs,
		}
		modelData = append(modelData, data)
	}
	res := model.DeviceInfoDB.WriteDB().Table(nclink_data.Namespace).Select(nclink_data.DataID, nclink_data.AdaptorID, nclink_data.DeviceID,
		nclink_data.ComponentID, nclink_data.DataItemID, nclink_data.Payload, nclink_data.AdaptorTime).Create(&modelData)
	if res.Error != nil {
		logger.Error("NCLinkSendData接口写入数据库失败 err:", res.Error)
		return nil, res.Error
	}
	logger.Infof("NCLinkSendData接口写入数据%v条", res.RowsAffected)
	return &nclink.NCLinkBaseResp{
		StatusCode: nclink.StatusOk,
		Detail:     "",
	}, nil
}
