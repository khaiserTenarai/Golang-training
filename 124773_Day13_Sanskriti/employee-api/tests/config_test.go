package tests

import (
	"testing"

	"employee-api/config"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_Defaults(t *testing.T) {
	t.Setenv("APP_ENV", "test")
	t.Setenv("PORT", "8081")

	cfg, err := config.LoadConfig("")
	assert.NoError(t, err)
	assert.Equal(t, "test", cfg.AppEnv)
	assert.Equal(t, "8081", cfg.Server.Port)
}
