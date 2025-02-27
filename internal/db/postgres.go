package db

import (
	"embed"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"link_service/internal/config"
)

//go:embed migrations/*.sql
var EmbedMigrations embed.FS


func NewPostgresDB(cfg config.Config) (*Queries, error) {
	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = MigrateDB(db); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(100)
	return New(db), err
}

func MigrateDB(db *sql.DB) error {
	goose.SetBaseFS(EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}
