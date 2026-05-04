package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LEVEL"  default:"debug"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewConfigMust() Config {
	var config Config
	if err := envconfig.Process("LOGGER", &config); err != nil {
		panic(fmt.Errorf(
			"process logger config: %w",
			err,
		))
	}
	return config
}
