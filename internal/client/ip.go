package client

import (
	"address-api/internal/config"
	"address-api/internal/metrics"
	v1 "address-api/pkg/api/proto/v1"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lockp111/go-easyzap"
	"github.com/pkg/errors"
	"github.com/sony/gobreaker"
)

type Ip struct {
	client  *gobreaker.CircuitBreaker
	url     string
	metrics *metrics.Metrics
}

func NewIpClient(cfg config.IpClientConfig, metrics *metrics.Metrics) *Ip {

	var st gobreaker.Settings
	st.Name = "Ip Client"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= uint32(cfg.Ip.MaxRetry) && failureRatio >= float64(cfg.Ip.MaxFailureRatio)
	}

	return &Ip{
		client:  gobreaker.NewCircuitBreaker(st),
		url:     cfg.Ip.URL,
		metrics: metrics,
	}
}

func (i *Ip) GetAddress(ctx context.Context, req *v1.IpRequest) (*v1.IpResponse, error) {

	body, err := i.client.Execute(func() (interface{}, error) {
		url := i.url + req.Ip
		resp, err := http.Get(url)
		if err != nil {
			mv := []string{"GetAddress", "error", "recover_address"}
			i.metrics.IpClient.WithLabelValues(mv...).Inc()
			errWrap := errors.Wrapf(err, "failed to recover address from ip %v", req.Ip)
			easyzap.Error(ctx, errWrap, "failed to recover address from ip")
			return nil, errWrap
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			mv := []string{"GetAddress", "error", "read_body"}
			i.metrics.IpClient.WithLabelValues(mv...).Inc()
			errWrap := errors.Wrapf(err, "failed to read body from ip %v", req.Ip)
			easyzap.Error(ctx, errWrap, "failed to read body from ip")
			return nil, errWrap
		}
		return body, nil
	})
	if err != nil {
		return nil, err
	}

	var resp *v1.IpResponse
	if err := json.Unmarshal(body.([]byte), &resp); err != nil {
		mv := []string{"GetAddress", "error", "unmarshal"}
		i.metrics.IpClient.WithLabelValues(mv...).Inc()
		errWrap := errors.Wrapf(err, "failed to unmarshal return api from ip %v", req.Ip)
		easyzap.Error(ctx, errWrap, "failed to unmarshal return api from ip")
		return nil, errWrap
	}

	if err := i.validateResponse(ctx, resp); err != nil {
		return nil, err
	}

	mv := []string{"GetAddress", "success", ""}
	i.metrics.IpClient.WithLabelValues(mv...).Inc()
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
			errWrap := errors.Wrapf(fmt.Errorf(v.Message), "invalid response ip client %v", resp.Status)
			easyzap.Error(ctx, errWrap, "invalid response response ip client")
			return errWrap
		}
	}

	return nil
}
