package config

import (
	"embed"
	"fmt"

	"github.com/lunarhue/libs-go/config"
	"github.com/lunarhue/libs-go/log"
)

//go:embed default.config.yaml
var defaultConfigFile embed.FS

type NixOSConfig struct {
	NixOSPath string `mapstructure:"nixos_path" description:"Path to the NixOS configuration file"`
}

type Config struct {
	DefaultPort int    `mapstructure:"default_port" description:"Port to listen on for incoming connections"`
	Mode        string `mapstructure:"mode" description:"Operation mode (server, agent, auto)"`
	K3sPath     string `mapstructure:"k3s_path" description:"Path to the K3s binary"`

	LogLevel string `mapstructure:"log_level" description:"Console logging level (debug, info, warn, error)"`
	LogFile  string `mapstructure:"log_file" description:"File to log to (empty for console only)"`
}

func Load() (*Config, error) {
	cfg, err := config.LoadConfig[Config](&defaultConfigFile, "default.config.yaml", "config.yaml", "METALLIC_FLOCK")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	log.SetLevelFromString(cfg.LogLevel)

	return cfg, nil
}
