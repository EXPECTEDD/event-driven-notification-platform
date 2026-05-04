package core_postgres_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	User     string        `envconfig:"USER"     required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Host     string        `envconfig:"HOST"     default:"localhost"`
	Port     string        `envconfig:"PORT"     default:"5432"`
	DBName   string        `envconfig:"DB"       required:"true"`
	SSLMode  string        `envconfig:"SSL"      default:"disable"`
	Timeout  time.Duration `envconfig:"TIMEOUT"  default:"30s"`
}

func NewConfigMust() Config {
	var config Config
	if err := envconfig.Process("POSTGRES", &config); err != nil {
		panic(fmt.Errorf("process connection pool config: %w",
			err,
		))
	}
	return config
}
