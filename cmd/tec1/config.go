//-----------------------------------------------------------------------------
/*

TEC-1A Configuration File

This is a toml (https://toml.io/en/) file
It stores the various parameters used to configure the tec-1a emulation.

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//-----------------------------------------------------------------------------

const configFile = "tec1a.cfg"

//-----------------------------------------------------------------------------

type array88Config struct {
	Enable bool `toml:"enable"` // is the 8x8 led array enabled?
}

type soundConfig struct {
	Enable bool `toml:"enable"` // is the sound enabled?
}

type Config struct {
	Array88 array88Config `toml:"array_8x8"`
	Sound   soundConfig   `toml:"sound"`
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
	return &Config{
		Sound: soundConfig{Enable: true},
	}
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
	return os.WriteFile(path, []byte(cfg.String()), 0664)
}

//-----------------------------------------------------------------------------
