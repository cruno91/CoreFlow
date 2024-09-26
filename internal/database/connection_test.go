package database

import (
	"testing"

	"CoreFlow/config"
	"github.com/spf13/viper"
)

func TestConnect(t *testing.T) {
	// Initialize Viper configuration
	config.InitConfig()

	// Set up database configuration using Viper
	cfg := Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.name"),
	}

	// Attempt to connect to the database
	err := Connect(cfg)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ensure that the connection can be pinged
	err = DB.Ping()
	if err != nil {
		t.Fatalf("Database connection is not alive: %v", err)
	}

	// Close the database connection after the test
	defer DB.Close()

	t.Log("Database connection test passed")
}
