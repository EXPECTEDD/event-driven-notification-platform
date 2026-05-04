package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host            string        `envconfig:"HOST"             required:"true"`
	Port            string        `envconfig:"PORT"             required:"true"`
	ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT"     default:"30s"`
	WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT"    default:"30s"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

func NewConfigMust() Config {
	var config Config
	if err := envconfig.Process("HTTP_SERVER", &config); err != nil {
		panic(fmt.Errorf(
			"process HTTP server config: %w",
			err,
		))
	}
	return config
}
