//-----------------------------------------------------------------------------
/*

Jupiter ACE Configuration File

This is a toml (https://toml.io/en/) file
It stores the various parameters used to configure the Jupiter ACE emulation.

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//-----------------------------------------------------------------------------

const configFile = "jace.cfg"

//-----------------------------------------------------------------------------

type soundConfig struct {
	Enable bool `toml:"enable"` // is the sound enabled?
}

type Config struct {
	Sound soundConfig `toml:"sound"`
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
