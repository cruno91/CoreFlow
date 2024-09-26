package main

import (
	"CoreFlow/cmd"
	"CoreFlow/config"
	"CoreFlow/internal/database"
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()

	dbConfig := database.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.name"),
	}

	err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			fmt.Printf("Could not close the database connection: %v", err)
		}
	}(database.DB)

	cmd.Execute()
}
