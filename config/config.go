package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Optionally, read environment variables
	viper.AutomaticEnv()
}
