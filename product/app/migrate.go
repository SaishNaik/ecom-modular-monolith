package app

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	_defaultAttempts     = 5
	_defaultRetryTimeout = time.Second
	_migrationFilePath   = "/db/"
)

func init() {
	pgUrl, ok := os.LookupEnv("PG_URL")
	if !ok && len(pgUrl) != 0 {
		log.Fatalf("PG_URL not available or empty in environment")
	}
	pgUrl += "?sslmode=disable"

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)
	for attempts > 0 {
		cur, _ := os.Getwd()
		dir := filepath.Dir(cur + "/../..")
		m, err = migrate.New(fmt.Sprintf("file://%s/%s", dir, _migrationFilePath), pgUrl)
		if err == nil {
			break
		}
		log.Printf("Postgres trying to connect, attempts remaining: %d", attempts-1)
		time.Sleep(_defaultRetryTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate Postgres connection error: %s", err)
	}

	defer m.Close()
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate up failed with error:%s", err)
		return
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("no migration needed")
		return
	}
	log.Println("migration done successfully")
}
