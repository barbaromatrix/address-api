package client

import (
	"address-api/internal/model"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/lockp111/go-easyzap"
	"github.com/sony/gobreaker"
)

type IIpClient interface {
	GetAddress(ctx context.Context, ae model.IpRequest) (*model.IpResponse, error)
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

func (a *IpClient) GetAddress(ctx context.Context, request model.IpRequest) (*model.IpResponse, error) {

	body, err := a.client.Execute(func() (interface{}, error) {
		url := a.url + request.Ip
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

	var IpResponse model.IpResponse
	if err := json.Unmarshal(body.([]byte), &IpResponse); err != nil {
		easyzap.Error(ctx, err, "[ip] failed during unmarshal return api")

		return nil, err
	}

	return &IpResponse, nil
}
