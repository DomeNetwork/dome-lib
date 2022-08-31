package sql

import (
	"fmt"

	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ SQL = &Postgres{}

// Postgres interacts with the Postgres3 local database.
type Postgres struct {
	cli *gorm.DB
}

// Close the connection.
func (db *Postgres) Close() (err error) {
	log.D("postgres", "close")
	return
}

// Create a new entry in the database.
func (db *Postgres) Create(v interface{}) (err error) {
	log.D("postgres", "create", v)
	err = db.cli.Create(v).Error
	return
}

// Delete from the database the object that matches the provided object.
func (db *Postgres) Delete(v interface{}) (err error) {
	log.D("postgres", "delete", v)
	err = db.cli.Delete(v).Error
	return
}

// Exec will allow for the execution of raw SQL statements.
func (db *Postgres) Exec(qry string) (err error) {
	log.D("postgres", "exec", qry)
	err = db.cli.Exec(qry).Error
	return
}

// Get will return an object that matches the provided interface.
func (db *Postgres) Get(v interface{}) (err error) {
	log.D("postgres", "get", v)
	err = db.cli.First(v).Error
	return
}

// GetWhere will return an object that matches the provided filter.
func (db *Postgres) GetWhere(filter map[string]interface{}, v interface{}) (err error) {
	log.D("postgres", "getWhere", filter, v)
	tx := db.cli.Limit(1)
	for key, value := range filter {
		tx.Where(key, value)
	}
	err = tx.First(v).Error
	return
}

// Migrate the database automatically by reading the provided interfaces.
func (db *Postgres) Migrate(v ...interface{}) (err error) {
	log.D("postgres", "migrate", v)
	vs := make([]interface{}, len(v))
	for i, o := range v {
		vs[i] = o
	}
	err = db.cli.AutoMigrate(vs...)
	return
}

// Open the connection.
func (db *Postgres) Open() (err error) {
	log.D("postgres", "open")
	ssl := "disable"
	if cfg.Bool("postgres.ssl") {
		ssl = "prefer"
	}
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s TimeZone=%s",
		cfg.Str("postgres.host"), cfg.Int("postgres.port"), cfg.Str("postgres.name"),
		cfg.Str("postgres.user"), cfg.Str("postgres.pass"), ssl, cfg.Str("postgres.tz"),
	)
	log.D("postgres", "open", dsn)
	db.cli, err = gorm.Open(postgres.Open(dsn))
	return
}

// Select entries from the database that match the query and options.
func (db *Postgres) Select(qry Query, opts Opts, vs interface{}) (err error) {
	log.D("postgres", "select", qry, opts, vs)
	tx := db.cli.Offset(opts.Offset).Limit(opts.Limit)
	for key, value := range qry {
		tx.Where(key, value)
	}
	err = tx.Find(vs).Error
	return
}

// Update the database for the matching entry of provided object.
func (db *Postgres) Update(v interface{}) (err error) {
	log.D("postgres", "update", v)
	err = db.cli.Save(v).Error
	return
}
