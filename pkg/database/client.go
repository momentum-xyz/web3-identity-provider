package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
)

func NewDB(cfg *Config) (*sql.DB, func(), error) {
	dsn := makeDSN(cfg)
	driver := cfg.Dialect
	if cfg.Dialect == "postgres" {
		driver = "pgx"
	}

	db, err := sql.Open(driver, dsn)

	if err != nil {
		return nil, nil, errors.Wrap(err, "")
	}

	if cfg.Migrate {
		err := Migrate(db, cfg)

		if err != nil {
			return nil, nil, err
		}
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}

func makeDSN(cfg *Config) string {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s", cfg.Dialect, cfg.Username, cfg.Password, cfg.Host, strconv.FormatUint(uint64(cfg.Port), 10), cfg.Name, cfg.DSNParams)
	return dsn
}
