package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/jvhab/Go-clean/config"
	"time"
)

const (
	setMaxOpenConns    = 60
	setConnMaxLifetime = 120
	setMaxIdleConns    = 30
	setConnMaxIdleTime = 20
)

func NewPostgresConn(cfg *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Login,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)
	db, err := sqlx.Connect(cfg.Postgres.PGDriver, dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(setMaxOpenConns)
	db.SetConnMaxLifetime(setConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(setMaxIdleConns)
	db.SetConnMaxIdleTime(setConnMaxIdleTime * time.Second)
	return db, nil
}
