package xorm

import (
	"context"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/database"
	"github.com/pkg/errors"
)

func NewEnt(db *sql.DB, cfg *database.Config) *ent.Client {
	return ent.NewClient(ent.Driver(entsql.OpenDB(cfg.Dialect, db)))
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
