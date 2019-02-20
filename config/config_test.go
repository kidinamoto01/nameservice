package config

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultInstrumentationConfig()
	require.True(t, cfg.Prometheus)
}
