package internal

import (
	"time"

	"github.com/caarlos0/env/v9"
)

type config struct {
	Base struct {
		StopKey []string `env:"HOOK_BASE_STOP_KEY" envDefault:"ctrl,shift,q"`
	}
	RepeatedClick struct {
		FireKey  []string      `env:"HOOK_FIRE_REPEATED_CLICKS_KEY" envDefault:"ctrl,shift,r,f"`
		StopKey  []string      `env:"HOOK_STOP_REPEATED_CLICKS_KEY" envDefault:"ctrl,shift,r,s"`
		Duration time.Duration `env:"REPEATED_CLICKS_INTERVAL" envDefault:"1s"`
		Interval time.Duration `env:"REPEATED_CLICKS_DURATION" envDefault:"1s"`
	}
}

func loadConfigFromEnviron() (config, error) {
	var c config
	err := env.Parse(&c)
	return c, err
}
