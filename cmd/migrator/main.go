package main

import (
	"log"

	"github.com/UnknownOlympus/db-migrator/internal/config"
	"github.com/UnknownOlympus/db-migrator/internal/repository"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func main() {
	cfg := config.MustLoad()

	dbpool, dbErr := repository.NewDatabase(
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name)
	if dbErr != nil {
		log.Fatalf("Failed to connect to DB: %v", dbErr)
	}

	dtb := stdlib.OpenDBFromPool(dbpool)
	if migrationErr := goose.Up(dtb, "migrations"); migrationErr != nil {
		log.Fatal(migrationErr)
	}
	defer dbpool.Close()

	log.Println("✅ Migrations applied successfully")
}
