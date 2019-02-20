package nservice

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	cfg "github.com/tendermint/tendermint/config"

)
const (
	MetricsSubsystem = "module_nameservice"
	SumLabel   = "SumNameService"
)
type Metrics struct {
	SumOfNameservice metrics.Gauge //

}


// PrometheusMetrics returns Metrics build using Prometheus client library.
func PrometheusMetrics(config *cfg.InstrumentationConfig) *Metrics {
	if !config.Prometheus {
		return NopMetrics()
	}
	return &Metrics{
		SumOfNameservice: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: config.Namespace,
			Subsystem: MetricsSubsystem,
			Name:      "sum_of_service",
			Help:      "the sum of nameservice",
		}, []string{}),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		SumOfNameservice: discard.NewGauge(),
	}
}

