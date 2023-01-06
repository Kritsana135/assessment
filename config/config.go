package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) {
	viper.SetDefault("GF_SHUTDOWN_TIMEOUT", 0)
	viper.SetDefault("AUTO_MIGRATE", false)
	viper.SetDefault("PORT", 2565)

	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Info("No configuration file found, using defaults")
	}
}
