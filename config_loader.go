package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

type Config struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

func loadConfig() (*Config, error) {
	v := viper.New()

	pflag.String("host", "defaultHost", "host")
	pflag.Int("port", 8080, "port")
	pflag.Bool("debug", false, "debug mode")
	pflag.Parse()

	// 1. Config file (lowest priority)
	v.SetConfigFile(defaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// 2. Environment variables (override config file)
	v.SetEnvPrefix("APP") // environment vars like APP_HOST
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 3. Command-line flags (highest priority)
	// Manually set if provided
	err := v.BindPFlags(pflag.CommandLine)
	if err != nil {
		return nil, fmt.Errorf("error binding flags: %w", err)
	}

	var cfg Config
	if err = v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}
	return &cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Output config
	fmt.Printf("Host:  %s\n", cfg.Host)
	fmt.Printf("Port:  %d\n", cfg.Port)
	fmt.Printf("Debug: %v\n", cfg.Debug)
}
