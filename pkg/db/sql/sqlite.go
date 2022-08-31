package sql

import (
	"github.com/domenetwork/dome-lib/pkg/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO: combine SQL interfaces, they all use GORM.

var _ SQL = &SQLite{}

// SQLite interacts with the SQLite3 local database.
type SQLite struct {
	cli *gorm.DB
}

// Close the connection.
func (db *SQLite) Close() (err error) {
	log.D("sqlite", "close")
	return
}

// Create a new entry in the database.
func (db *SQLite) Create(v interface{}) (err error) {
	log.D("sqlite", "create", v)
	err = db.cli.Create(v).Error
	return
}

// Delete from the database the object that matches the provided object.
func (db *SQLite) Delete(v interface{}) (err error) {
	log.D("sqlite", "delete", v)
	err = db.cli.Delete(v).Error
	return
}

// Exec will allow for the execution of raw SQL statements.
func (db *SQLite) Exec(qry string) (err error) {
	log.D("sqlite", "exec", qry)
	err = db.cli.Exec(qry).Error
	return
}

// Get will return an object that matches the provided interface.
func (db *SQLite) Get(v interface{}) (err error) {
	log.D("sqlite", "get", v)
	err = db.cli.First(v).Error
	return
}

// GetWhere will return an object that matches the provided filter.
func (db *SQLite) GetWhere(filter map[string]interface{}, v interface{}) (err error) {
	log.D("sqlite", "getWhere", filter, v)
	tx := db.cli.Limit(1)
	for key, value := range filter {
		tx.Where(key, value)
	}
	err = tx.First(v).Error
	return
}

// Migrate the database automatically by reading the provided interfaces.
func (db *SQLite) Migrate(v ...interface{}) (err error) {
	log.D("sqlite", "migrate", v)
	vs := make([]interface{}, len(v))
	for i, o := range v {
		vs[i] = o
	}
	err = db.cli.AutoMigrate(vs...)
	return
}

// Open the connection.
func (db *SQLite) Open() (err error) {
	log.D("sqlite", "open")
	db.cli, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	return
}

// Select entries from the database that match the query and options.
func (db *SQLite) Select(qry Query, opts Opts, vs interface{}) (err error) {
	log.D("sqlite", "select", qry, opts, vs)
	tx := db.cli.Offset(opts.Offset).Limit(opts.Limit)
	for key, value := range qry {
		tx.Where(key, value)
	}
	err = tx.Find(vs).Error
	return
}

// Update the database for the matching entry of provided object.
func (db *SQLite) Update(v interface{}) (err error) {
	log.D("sqlite", "update", v)
	err = db.cli.Save(v).Error
	return
}
