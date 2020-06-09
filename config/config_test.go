package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	FilePath = "config.development.yml"
)

func TestNewConfigShouldReturnConfig(t *testing.T) {
	opt := &Options{FilePath: FilePath}
	cfg := New(opt)
	require.NotNil(t, cfg)
	require.Equal(t, "commons", cfg.ApiKey)
}
