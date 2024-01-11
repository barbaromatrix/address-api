package service

import (
	"address-api/internal/client"
	"address-api/internal/model"
	"context"
)

type Ip struct {
	client client.IIpClient
}

func NewIpService(client client.IIpClient) *Ip {
	return &Ip{
		client: client,
	}
}

func (i *Ip) GetFormBy(ctx context.Context, request model.IpRequest) (*model.IpResponse, error) {

	forms, err := i.client.GetAddress(ctx, request)
	if err != nil {
		return nil, err
	}

	return forms, nil
}
