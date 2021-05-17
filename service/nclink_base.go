package service

import (
	"context"

	"github.com/mk1010/industry_adaptor/nclink"
)

type NcLinkServiceProvider struct {
	*nclink.NCLinkServiceProviderBase
}

func Init(ctx context.Context) error {
	return nil
}

func NewNcLinkServiceProvider() *NcLinkServiceProvider {
	return &NcLinkServiceProvider{
		&nclink.NCLinkServiceProviderBase{},
	}
}
