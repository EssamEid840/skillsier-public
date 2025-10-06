package database

import (
	"database/sql"
	"users.go/m/internal/infrastructure/config"
	"users.go/m/internal/infrastructure/database/postgres"
)

type DatabaseProvider string

const (
	Postgres DatabaseProvider = "postgres"
)

type Database interface {
	GetConnection(connection string) (*sql.DB, error)
	CloseAll()
}

func New(provider DatabaseProvider, config config.Config) Database {
	switch provider {
	case Postgres:
		return postgres.New(config)
	default:
		return nil
	}

}
