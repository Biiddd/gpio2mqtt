package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	MQTT MQTTConfig `mapstructure:"mqtt"`
	Log  LogConfig  `mapstructure:"log"`
	App  AppConfig  `mapstructure:"app"`
}

type MQTTConfig struct {
	Broker   string `mapstructure:"broker"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	ClientID string `mapstructure:"client_id"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	File       string `mapstructure:"file"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type AppConfig struct {
	Di map[string]DiConfig `mapstructure:"di"`
	Do map[string]DoConfig `mapstructure:"do"`
}

type DiConfig struct {
	Path        string `mapstructure:"path"`
	Interval    int    `mapstructure:"interval"`
	StatusTopic string `mapstructure:"status-topic"`
}

type DoConfig struct {
	Path        string `mapstructure:"path"`
	Interval    int    `mapstructure:"interval"`
	StatusTopic string `mapstructure:"status-topic"`
	CmdTopic    string `mapstructure:"cmd-topic"`
	HighPayload string `mapstructure:"high-payload"`
	LowPayload  string `mapstructure:"low-payload"`
}

func initConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	// 解析到结构体
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("parse config error: %w", err))
	}

	err = configCheck()
	if err != nil {
		panic(fmt.Errorf("parse config error: %w", err))
	}

	log.Printf("config load success: %+v\n", config)
}

func configCheck() error {
	for name, di := range config.App.Di {
		if di.Path == "" {
			return fmt.Errorf("DI[%s] config error: path can't be null", name)
		}
		if di.StatusTopic == "" {
			return fmt.Errorf("DI[%s] config error: status-topic can't be null", name)
		}
		if di.Interval <= 0 {
			return fmt.Errorf("DI[%s] config error: interval must > 0", name)
		}
	}

	for name, do := range config.App.Do {
		if do.Path == "" {
			return fmt.Errorf("DO[%s] config error: path can't be null", name)
		}
		if do.StatusTopic == "" {
			return fmt.Errorf("DO[%s] config error: status-topic can't be null", name)
		}
		if do.CmdTopic == "" {
			return fmt.Errorf("DO[%s] config error: cmd-topic can't be null", name)
		}

		if do.Interval <= 0 {
			return fmt.Errorf("DO[%s] config error: interval must > 0", name)
		}

		if do.HighPayload == "" && do.LowPayload == "" {
			return fmt.Errorf("DO[%s] config error: high-payload & low-payload can't be null both", name)
		}
	}

	return nil
}
