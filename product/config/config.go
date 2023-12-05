package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Postgres Postgres `mapstructure:"postgres"`
}

// todo add other params
type Postgres struct {
	DSN string `mapstructure:"dsn"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	//get curr dir
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	//todo load config from env rather than file in code
	viper.SetConfigName("product-config")
	viper.AddConfigPath(dir + "/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}
