package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfigFile(t *testing.T) {
	c := ParseConfigFile("config.yml.example")
	assert.Equal(t, c.Binance.ApiKey, "abc")
	assert.Equal(t, c.Binance.SecretKey, "123")
}
