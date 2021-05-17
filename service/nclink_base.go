package service

import (
	"context"

	"github.com/apache/dubbo-go/config"
	"github.com/mk1010/industry_adaptor/nclink"
)

type NcLinkServiceProvider struct {
	*nclink.NCLinkServiceProviderBase
}

func Init(ctx context.Context) error {
	config.SetProviderService(NewNcLinkServiceProvider())
	return nil
}

func NewNcLinkServiceProvider() *NcLinkServiceProvider {
	return &NcLinkServiceProvider{
		&nclink.NCLinkServiceProviderBase{},
	}
}

func (u *NcLinkServiceProvider) Reference() string {
	return "nCLinkServiceImpl"
}
