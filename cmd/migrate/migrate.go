package migrate

import (
	"CoreFlow/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize configuration
		config.InitConfig()

		//Build the database URL
		dbURL := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			viper.GetString("database.user"),
			viper.GetString("database.password"),
			viper.GetString("database.host"),
			viper.GetInt("database.port"),
			viper.GetString("database.name"),
		)

		//Run migrations
		if err := runMigrations(dbURL); err != nil {
			log.Fatalf("Could not run migrations: %v", err)
		}

		fmt.Println("Migrations ran successfully")
	},
}

func runMigrations(dbURL string) error {
	// Initialize the migrate instance
	m, err := migrate.New(
		"file://internal/database/migrations", // Path to the migrations directory
		dbURL,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate: %w", err)
	}

	// Run migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
