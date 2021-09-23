package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/sqoopdata/madoc/internal/config"
)

func main() {
	godotenv.Load()

	cfg := config.Get()

	direction := cfg.GetMigration()
	if direction != "down" && direction != "up" {
		log.Println("-migrate accepts [up, down] values only")
		return
	}

	m, err := migrate.New("file://db/migrations", cfg.GetDBConnStr())
	if err != nil {
		log.Printf("%s", err)
		return
	}

	if direction == "up" {
		if err := m.Up(); err != nil {
			log.Printf("failed migrate up: %s", err)
			return
		}
	}

	if direction == "down" {
		if err := m.Down(); err != nil {
			log.Printf("failed migrate down: %s", err)
			return
		}
	}
}
