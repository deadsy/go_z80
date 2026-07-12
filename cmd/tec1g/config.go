//-----------------------------------------------------------------------------
/*

TEC-1G Configuration File

This is a toml (https://toml.io/en/) file
It stores the various parameters used to configure the tec-1g emulation.

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/deadsy/go_z80/device/rtc"
)

//-----------------------------------------------------------------------------

const configFile = "tec1g.cfg"

//-----------------------------------------------------------------------------

type Config struct {
	RTC rtc.Config `toml:"rtc"`
}

func (cfg *Config) String() string {
	s, err := toml.Marshal(cfg)
	if err != nil {
		log.Printf("unable to marshal config structure: %s", err)
		return ""
	}
	return string(s)
}

//-----------------------------------------------------------------------------

// return a default config
func defaultConfig() *Config {
	// rtc
	rtc := rtc.Config{
		Enable: true,
	}
	cfg := &Config{
		RTC: rtc,
	}
	return cfg
}

// load the config from a file
func loadConfig(path string) (*Config, error) {
	s, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = toml.Unmarshal(s, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// save the config to a file
func (cfg *Config) saveConfig(sys *system, path string) error {
	// rtc may have changed
	cfg.RTC = sys.rtc.GetConfig()
	return os.WriteFile(path, []byte(cfg.String()), 0664)
}

//-----------------------------------------------------------------------------
