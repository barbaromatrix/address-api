package client

import (
	"address-api/internal/model"
	v1 "address-api/pkg/api/proto/v1"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/lockp111/go-easyzap"
	"github.com/sony/gobreaker"
)

type IIpClient interface {
	GetAddress(ctx context.Context, req *v1.IpRequest) (*v1.IpResponse, error)
}

type IpClient struct {
	client *gobreaker.CircuitBreaker
	url    string
}

func NewIIpClient(cfg model.Config) *IpClient {

	var st gobreaker.Settings
	st.Name = "Ip Client"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= uint32(cfg.Ip.HTTP.MaxRetry) && failureRatio >= float64(cfg.Ip.HTTP.MaxFailureRatio)
	}

	return &IpClient{
		client: gobreaker.NewCircuitBreaker(st),
		url:    cfg.Ip.HTTP.URL,
	}
}

func (a *IpClient) GetAddress(ctx context.Context, req *v1.IpRequest) (*v1.IpResponse, error) {

	body, err := a.client.Execute(func() (interface{}, error) {
		url := a.url + req.Ip
		resp, err := http.Get(url)
		if err != nil {
			easyzap.Error(ctx, err, "error to recover address from ip")

			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			easyzap.Error(ctx, err, "[ip] failed during read body")

			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	var resp *v1.IpResponse
	if err := json.Unmarshal(body.([]byte), &resp); err != nil {
		easyzap.Error(ctx, err, "[ip] failed during unmarshal return api")

		return nil, err
	}

	return resp, nil
}
