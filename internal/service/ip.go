package service

import (
	v1 "address-api/pkg/api/proto/v1"
	"context"
	"fmt"

	"github.com/lockp111/go-easyzap"
	"github.com/pkg/errors"
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

func (i *Ip) GetAddress(ctx context.Context, request *v1.IpRequest) (*v1.IpResponse, error) {

	resp, err := i.client.GetAddress(ctx, request)
	if err != nil {

		return nil, err
	}

	if resp.Status != "success" {

		return nil, errors.Wrap(err, "Failed get address from ip")
	}

	if i.validateResponse(ctx, resp) != nil {

		return nil, errors.Wrap(err, "Status response from ip is not success")
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
