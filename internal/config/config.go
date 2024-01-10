package config

import (
	"address-api/internal/model"
	"context"
	"os"
	"sync"

	"github.com/integralist/go-findroot/find"
	"github.com/lockp111/go-easyzap"
	"github.com/spf13/viper"
)

const (
	ENV_PROFILE_LOCAL = "local"
)

var (
	runOnce sync.Once
	config  model.Config
)

func initConfig() {
	envProfile := os.Getenv("ENV_PROFILE")
	if envProfile == ENV_PROFILE_LOCAL {
		setEnvsByFile()
	}
	viper.AutomaticEnv()
}

func setEnvsByFile() {
	root, _ := find.Repo()

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(root.Path + "/build/package/env/local")

	if err := viper.ReadInConfig(); err != nil {
		easyzap.Panic(context.Background(), err, "failed reading config file")
	}
}

func GetConfig() model.Config {
	runOnce.Do(func() {
		initConfig()
		config = model.Config{
			AppName:      viper.GetString("APPLICATION_NAME"),
			ServerPort:   viper.GetString("SERVER_PORT"),
			HealthPort:   viper.GetString("HEALTH_PORT"),
			TimeLocation: viper.GetString("TIME_LOCATION"),
			ViaCep: model.HTTPClientConfig{
				HTTP: model.HTTP{
					URL:      viper.GetString("VIA_CEP_URL"),
					MaxRetry: viper.GetInt("VIA_CEP_MAX_RETRY"),
					Timeout:  viper.GetInt("VIA_CEP_MAX_TIMEOUT"),
				},
			},
		}
	})
	return config
}
