package service

import (
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
