package db

import (
	"backend/internal/profile/infrastructure/db/migrations"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
)

func Migrate(db *sql.DB, targetVersion int) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	s := bindata.Resource(migrations.AssetNames(), func(name string) ([]byte, error) {
		return migrations.Asset(name)
	})
	d, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("go-bindata", d, "mysql", driver)
	if err != nil {
		return err
	}
	if targetVersion > 0 {
		return m.Force(targetVersion)
	}
	return m.Up()
}
