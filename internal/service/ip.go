package service

import (
	"address-api/internal/client"
	v1 "address-api/pkg/api/proto/v1"
	"context"
	"fmt"

	"github.com/lockp111/go-easyzap"
)

type IIp interface {
	GetAddress(ctx context.Context, request *v1.IpRequest) (*v1.IpResponse, error)
}

type Ip struct {
	client client.IIpClient
}

func NewIpService(client client.IIpClient) *Ip {
	return &Ip{
		client: client,
	}
}

func (i *Ip) GetAddress(ctx context.Context, request *v1.IpRequest) (*v1.IpResponse, error) {

	resp, err := i.client.GetAddress(ctx, request)
	if err != nil {

		return nil, err
	}

	if resp.Status != "success" {

		return nil, err
	}

	if i.validateResponse(ctx, resp) != nil {

		return nil, err
	}

	return resp, nil
}

func (i *Ip) validateResponse(ctx context.Context, resp *v1.IpResponse) error {
	validations := []struct {
		Condition bool
		Message   string
	}{
		{resp.Status != "success", "ip-api fail"},
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
