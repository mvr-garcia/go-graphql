package app

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	Database struct {
		Driver string
		DSN    string
	}
}

func LoadConfig() *Config {

	viper.SetDefault("Port", "8080")
	viper.SetDefault("Database.Driver", "sqlite3")
	viper.SetDefault("Database.DSN", "./data.db")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
