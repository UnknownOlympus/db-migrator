package config_test

import (
	"testing"

	"github.com/Flaque/filet"
	"github.com/UnknownOlympus/db-migrator/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_MustLoadFromFile(t *testing.T) {
	envContent := `
DATABASE_HOST=host
DATABASE_PORT=3000
DATABASE_USERNAME=admin
DATABASE_PASSWORD=adminpass
DATABASE_NAME=test
`
	filet.File(t, ".env", envContent)
	defer filet.CleanUp(t)

	cfg := config.MustLoad()

	assert.Equal(t, "host", cfg.Host)
	assert.Equal(t, "3000", cfg.Port)
	assert.Equal(t, "admin", cfg.Username)
	assert.Equal(t, "adminpass", cfg.Password)
	assert.Equal(t, "test", cfg.Name)
}
