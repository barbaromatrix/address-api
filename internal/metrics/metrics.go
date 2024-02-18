package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	onceConfig    sync.Once
	metricConfigs = &Metrics{}
)

type Metrics struct {
	IpClient *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	onceConfig.Do(func() {
		metricConfigs.IpClient = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "ip_client",
			Help: "Return latitude and longitude based on IP address",
		}, []string{"method", "status", "details"})
		prometheus.MustRegister(metricConfigs.IpClient)
	})
	return metricConfigs
}
