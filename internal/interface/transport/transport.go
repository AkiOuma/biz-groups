package transport

import (
	v1 "github.com/AkiOuma/biz-groups/api/goapi/janus-biz-groups/service/v1"
	"github.com/AkiOuma/biz-groups/internal/usecase"
)

type Transport struct {
	uc usecase.Usecase
	v1.UnimplementedTransportServer
}

var _ v1.TransportServer = (*Transport)(nil)

func NewTransport(uc usecase.Usecase) *Transport {
	return &Transport{
		uc: uc,
	}
}
