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

type IAwesomeClient interface {
	GetAddress(ctx context.Context, ae model.AweSomeRequest) (*model.AweSomeResponse, error)
}

type AwesomeClient struct {
	client *gobreaker.CircuitBreaker
	url    string
}

func NewAwesomeClient(cfg model.Config) *AwesomeClient {

	var st gobreaker.Settings
	st.Name = "Awesome Client"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= uint32(cfg.AweSome.HTTP.MaxRetry) && failureRatio >= float64(cfg.AweSome.HTTP.MaxFailureRatio)
	}

	return &AwesomeClient{
		client: gobreaker.NewCircuitBreaker(st),
		url:    cfg.AweSome.HTTP.URL,
	}
}

func (a *AwesomeClient) GetAddress(ctx context.Context, cep model.AweSomeRequest) (*model.AweSomeResponse, error) {

	body, err := a.client.Execute(func() (interface{}, error) {
		url := a.url + cep.Cep
		resp, err := http.Get(url)
		if err != nil {
			easyzap.Error(context.Background(), err, "error to recover address from ads awesome")

			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			easyzap.Error(context.Background(), err, "[awesome] failed during read body")

			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	var aweSomeResponse model.AweSomeResponse
	if err := json.Unmarshal(body.([]byte), &aweSomeResponse); err != nil {
		easyzap.Error(context.Background(), err, "[awesome] failed during unmarshal return api")

		return nil, err
	}

	return &aweSomeResponse, nil
}
