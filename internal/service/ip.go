package service

import (
	v1 "address-api/pkg/api/proto/v1"
	"context"
)

type IpClient interface {
	GetAddress(ctx context.Context, request *v1.IpRequest) (*v1.IpResponse, error)
}

type Ip struct {
	client IpClient
}

func NewIpService(client IpClient) *Ip {
	return &Ip{
		client: client,
	}
}

func (i *Ip) GetAddress(ctx context.Context, req *v1.IpRequest) (*v1.IpResponse, error) {

	resp, err := i.client.GetAddress(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
