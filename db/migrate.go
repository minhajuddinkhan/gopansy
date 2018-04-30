package db

import (
	"database/sql"

	"github.com/DavidHuie/gomigrate"

	_ "github.com/lib/pq"
	conf "github.com/minhajuddinkhan/gopansy/config"
	"github.com/minhajuddinkhan/gopansy/constants"
)

//Migrator Migrator
type Migrator struct {
	Conf conf.Configuration
}

//Migrate Migrates the database
func (m *Migrator) Migrate(path string) error {

	db, err := sql.Open(constants.DbType, m.Conf.ConnectionString)
	if err != nil {
		return err
	}
	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, path)
	return migrator.Migrate()
}
