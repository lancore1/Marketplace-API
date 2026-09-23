package main

import (
	"context"
	"fmt"
	"log"
	"module/internal/config"
	"module/internal/database"

	"github.com/jackc/pgx/v5"
)

const (
	MigrationPath = "migration"
	DriverName    = "postgres"
)

var ctx = context.Background()

func main() {
	//Load CFG
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to DB
	conn, err := pgx.Connect(ctx, cfg.DB.URL)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(ctx)

	var migrator = database.Migration{}
	if err := migrator.RunMigration(MigrationPath, DriverName, cfg.DB.URL); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Migration is succesfull")
	}

}
