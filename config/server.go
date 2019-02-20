package config

import(
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"

)

// server context
type Context struct {
	Config *cfg.InstrumentationConfig
	Logger log.Logger
}


// DefaultInstrumentationConfig returns a default configuration for metrics
// reporting.
func DefaultInstrumentationConfig() *cfg.InstrumentationConfig {
	return &cfg.InstrumentationConfig{
		Prometheus:           true,
		PrometheusListenAddr: ":26660",
		MaxOpenConnections:   3,
		Namespace:            "tendermint",
	}
}


func (c *Context) SetEnablePrometheus(isEnable bool) {
	c.Config.Prometheus = isEnable
}

// GetMinGasPrices returns the validator's minimum gas prices based on the set
// configuration.
func (c *Context) GetEnablePrometheus() bool {

	return 	c.Config.Prometheus
}

