package service

import (
	"context"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/mk1010/industry_adaptor/nclink"
)

func (n *NcLinkServiceProvider) NCLinkAuth(ctx context.Context, req *nclink.NCLinkAuthReq) (*nclink.NCLinkAuthResp, error) {
	logger.Info("NCLinkAuth接口收到请求:", req)
	return &nclink.NCLinkAuthResp{
		BaseResp: &nclink.NCLinkBaseResp{
			StatusCode: nclink.StatusOk,
			Detail:     "auth 未启用",
		},
	}, nil
}
