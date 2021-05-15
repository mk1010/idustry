package service

import (
	"context"
	"testing"

	"github.com/mk1010/industry_adaptor/nclink"
)

func TestGetMeta(t *testing.T) {
	b := NewNcLinkServiceProvider()
	resp, err := b.NCLinkGetMeta(context.Background(), &nclink.NCLinkMetaDataReq{
		AdaptorId: []string{"123"},
	})
	t.Log(resp, err)
}
