package db

import (
	"database/sql"

	conf "github.com/minhajuddinkhan/gopansy/config"

	"github.com/minhajuddinkhan/gopansy/constants"
)

//Seeder Seeder
type Seeder struct {
	Conf conf.Configuration
}

// Seed  Seed
func (s *Seeder) Seed() error {

	db, err := sql.Open(constants.DbType, s.Conf.ConnectionString)
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO roles (name) SELECT 'admin' WHERE NOT EXISTS ( SELECT role_id from roles WHERE role_id = 1)")
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO users
			(username, hashed_password, role_id, permit_one_allowed, permit_two_allowed)
			 SELECT 'pancyAdmin', '$2a$05$ZW7dtscHYyl0B7OUlHJ4oOfsJsVt1adavPbpvXi5OjydxM4Tc3QFW', 1, true, true 
			 WHERE NOT EXISTS (SELECT user_id from users where user_id = 1)`)

	return err

}

//UnSeed UnSeed
func (s *Seeder) UnSeed() error {

	db, err := sql.Open(constants.DbType, s.Conf.ConnectionString)
	defer db.Close()

	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE from users where id = 1")
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE from roles where id = 1")

	return err

}
