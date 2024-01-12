package api

import (
	"address-api/internal/service"
	v1 "address-api/pkg/api/proto/v1"
)

type AddressV1 struct {
	v1.UnimplementedAddressServer
	ipService service.IIp
}

func NewAddressV5(ipService service.IIp) v1.AddressServer {
	return &AddressV1{
		ipService: ipService,
	}
}
