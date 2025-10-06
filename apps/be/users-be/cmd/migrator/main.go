// cmd/migrator/main.go
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // file:// source
)

func main() {
	// ---- Flags / Config ----
	// Default DSN matches your docker-compose (Postgres exposed on localhost)
	dsn := getenv("DB_DSN", "postgres://user:pass@localhost:5432/users?sslmode=disable")

	dir := flag.String("dir", "./migrations", "migrations directory")
	action := flag.String("action", "up", "up | down | steps | force | version | drop")
	steps := flag.Int("n", 1, "steps for -action=steps (negative to go down)")
	forceVersion := flag.Int("version", 0, "version for -action=force")
	flag.Parse()

	src := "file://" + mustAbs(*dir)

	// ---- Init migrator ----
	m, err := migrate.New(src, dsn)
	if err != nil {
		log.Fatalf("migrate.New: %v", err)
	}
	// Close returns two errors; handle both
	defer func() {
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			log.Printf("migrate close: sourceErr=%v databaseErr=%v", srcErr, dbErr)
		}
	}()

	// ---- Execute action ----
	switch *action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migrate up: %v", err)
		}
		log.Println("migrate up: OK (or no change)")

	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("migrate down 1: %v", err)
		}
		log.Println("migrate down 1: OK")

	case "steps":
		if *steps == 0 {
			log.Fatalf("-action=steps requires -n with non-zero steps (positive or negative)")
		}
		if err := m.Steps(*steps); err != nil {
			log.Fatalf("migrate steps %d: %v", *steps, err)
		}
		log.Printf("migrate steps %d: OK\n", *steps)

	case "force":
		if *forceVersion == 0 {
			log.Fatalf("-action=force requires -version=<number>")
		}
		if err := m.Force(*forceVersion); err != nil {
			log.Fatalf("migrate force %d: %v", *forceVersion, err)
		}
		log.Printf("migrate force %d: OK\n", *forceVersion)

	case "version":
		v, dirty, err := m.Version()
		if err != nil && err != migrate.ErrNilVersion {
			log.Fatalf("migrate version: %v", err)
		}
		if err == migrate.ErrNilVersion {
			log.Println("version: none (database empty); dirty=false")
		} else {
			log.Printf("version: %d; dirty=%v\n", v, dirty)
		}

	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatalf("migrate drop: %v", err)
		}
		log.Println("migrate drop: OK")

	default:
		log.Fatalf("unknown -action %q (use: up | down | steps | force | version | drop)", *action)
	}
}

// getenv returns env var k or def if empty.
func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

// mustAbs returns absolute path or exits on error.
func mustAbs(p string) string {
	a, err := filepath.Abs(p)
	if err != nil {
		log.Fatalf("abs(%q): %v", p, err)
	}
	return a
}
