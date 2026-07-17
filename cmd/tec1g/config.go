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
	"github.com/deadsy/go_z80/device/ds1302"
)

//-----------------------------------------------------------------------------

const configFile = "tec1g.cfg"

//-----------------------------------------------------------------------------

type array88Config struct {
	Enable bool `toml:"enable"` // is the 8x8 led array enabled?
}

type Config struct {
	RTC     ds1302.Config `toml:"rtc"`
	DIP     dipSwitch     `toml:"dip_switch"`
	Array88 array88Config `toml:"array_8x8"`
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
	rtc := ds1302.Config{
		Enable:        true,
		BaseYear:      2000,
		WeekDayOffset: 6,
	}
	dip := dipSwitch{}
	cfg := &Config{
		RTC: rtc,
		DIP: dip,
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
	cfg.RTC = sys.io.dev.rtc.GetConfig()
	return os.WriteFile(path, []byte(cfg.String()), 0664)
}

//-----------------------------------------------------------------------------
