package config

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type RTC struct {
	Enable bool          `toml:"enable"`
	RAM    [31]byte      `toml:"ram"`
	Offset time.Duration `toml:"offset"`
}

type Config struct {
	RTC RTC `toml:"rtc"`
}

func (cfg *Config) Save(path string) error {
	s, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", s)
	return nil
}
