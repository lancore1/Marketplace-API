package main

import (
	"context"
	"fmt"
	"log"
	"module/internal/config"
	"module/internal/database"
	"sync"
	"time"
)

const (
	MigrationPath = "migration"
	DriverName    = "postgres"
)

var ctx = context.Background()
var wg = sync.WaitGroup{}
var rw = sync.RWMutex{}

var poolCfg = database.PoolConfig{
	MaxConns:        25,               // Maximum number of connections in the pool
	MinConns:        5,                // Minimum number of connections (reserve)
	MaxConnLifetime: 1 * time.Hour,    // Maximum connection lifetime
	MaxConnIdleTime: 30 * time.Minute, // Maximum connection idle time before closing
}

func main() {
	//Load CFG
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to DB
	ctxTimeOut, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	connection, err := database.ConnectPGX(ctxTimeOut, cfg.DB.URL, &poolCfg)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	} else {
		fmt.Println("Connect succesfull")
	}

	defer connection.Close()
}
