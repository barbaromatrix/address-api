package api

import (
	"address-api/internal/service"
	v1 "address-api/pkg/api/proto/v1"
	"context"
	"fmt"

	"github.com/lockp111/go-easyzap"
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

func (a *AddressV1) GetAddressByIp(ctx context.Context, req *v1.IpRequest) (*v1.IpResponse, error) {
	if err := a.validateRequest(ctx, req); err != nil {

		return nil, err
	}

	resp, err := a.ipService.GetAddress(ctx, req)
	if err != nil {

		return nil, err
	}

	return resp, nil
}

func (a *AddressV1) validateRequest(ctx context.Context, req *v1.IpRequest) error {
	validations := []struct {
		Condition bool
		Message   string
	}{
		{req.Ip == "", "ip is empty"},
	}

	for _, v := range validations {
		if v.Condition {
			err := fmt.Errorf(v.Message)
			easyzap.Error(ctx, err, v.Message)
			return err
		}
	}

	return nil
}
