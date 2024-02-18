package config

import (
	"context"
	"strings"

	"github.com/joomcode/errorx"
	"github.com/lockp111/go-easyzap"
	"github.com/spf13/viper"
)

type Config struct {
	AppName      string
	ServerHost   string
	HealthPort   string
	TimeLocation string
}

type IpClientConfig struct {
	Ip HTTP
}

type HTTP struct {
	URL             string
	MaxRetry        int
	MaxFailureRatio int
	Name            string
}

var config = viper.New()

func Init() {
	config.AddConfigPath("internal/config/")
	config.SetConfigName("configuration")
	config.SetConfigType("yml")

	setConfigDefaults()

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			err := errorx.Decorate(err, "Error reading config file: %s", err)
			easyzap.Fatal(context.Background(), err, "Unable to keep the service without config file")
		}
	}

	config.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	config.AutomaticEnv()
}

func setConfigDefaults() {
	config.SetDefault("app.name", "address-api")
	config.SetDefault("server.port", "0.0.0.0:8080")
	config.SetDefault("health.port", "0.0.0.0:8081")
	config.SetDefault("timeLocation", "America/Sao_Paulo")
	config.SetDefault("ip.url", "http://ip-api.com/json/")
	config.SetDefault("ip.maxRetry", 2)
	config.SetDefault("ip.maxFailureRatio", 0.6)
	config.SetDefault("ip.name", "HTTP GET")
}

func GetIpClientConfig() IpClientConfig {
	return IpClientConfig{
		Ip: HTTP{
			URL:             config.GetString("ip.url"),
			MaxRetry:        config.GetInt("ip.maxRetry"),
			MaxFailureRatio: config.GetInt("ip.maxFailureRatio"),
			Name:            config.GetString("ip.name"),
		},
	}
}

func GetConfig() Config {
	return Config{
		AppName:      config.GetString("app.name"),
		ServerHost:   config.GetString("server.port"),
		HealthPort:   config.GetString("health.port"),
		TimeLocation: config.GetString("timeLocation"),
	}
}
