package mysql

import (
	"context"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/log"

	"github.com/pkg/errors"
)

func Migrate(db *sql.DB) error {
	log.Logln(0, "Migrating the database...")
	ent := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
	err := ent.Schema.Create(context.Background())

	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
