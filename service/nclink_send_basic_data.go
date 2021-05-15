package service

import (
	"context"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/mk1010/industry_adaptor/nclink"
)

func (n *NcLinkServiceProvider) NCLinkSendBasicData(ctx context.Context, req *nclink.NCLinkTopicMessage) (*nclink.NCLinkBaseResp, error) {
	logger.Info("NCLinkSendBasicData接口收到请求:", req)
	return &nclink.NCLinkBaseResp{
		StatusCode: nclink.StatusOk,
		Detail:     "SendBasicData 未启用",
	}, nil
}
