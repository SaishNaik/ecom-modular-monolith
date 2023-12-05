package postgresql

import (
	"context"
	"database/sql"
	"github.com/SaishNaik/ecom-modular-monolith/product/domain"
	"log"
	"log/slog"
	"time"
)

const (
	_defaultConnAttempts     = 3
	_defaultConnRetryTimeout = time.Second
)

type Postgres struct {
	db               *sql.DB
	connAttempts     int
	connRetryTimeout time.Duration
}

var _ domain.IProductRepo = (*Postgres)(nil)

func NewPostgres(dsn string) (domain.IProductRepo, error) {
	slog.Info("Postgres URL", dsn)
	pg := Postgres{
		connAttempts:     _defaultConnAttempts,
		connRetryTimeout: _defaultConnRetryTimeout,
	}

	var err error
	for pg.connAttempts > 0 {
		pg.db, err = sql.Open("postgres", dsn)
		if err == nil {
			break
		}
		log.Printf("Postgres is trying to connect,attempts left: %d", pg.connAttempts)
		time.Sleep(pg.connRetryTimeout)
		pg.connAttempts--
	}

	if err != nil {
		return nil, err
	}
	slog.Info("Connected to Postgres")
	return &pg, err
}

func (pg *Postgres) GetAllProducts(context.Context) ([]*domain.ProductModel, error) {
	return nil, nil
}
