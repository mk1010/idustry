package handler

import (
	"context"

	"github.com/mk1010/industry_adaptor/nclink"
)

type NcLinkServiceProvider struct {
	*nclink.NCLinkServiceProviderBase
}

func NewNcLinkServiceProvider() *NcLinkServiceProvider {
	return &NcLinkServiceProvider{
		&nclink.NCLinkServiceProviderBase{},
	}
}

func (n *NcLinkServiceProvider) NCLinkGetMeta(ctx context.Context, req *nclink.NCLinkMetaDataReq) (*nclink.NCLinkMetaDataResp, error) {
	resp := new(nclink.NCLinkMetaDataResp)
	return resp, nil
}
