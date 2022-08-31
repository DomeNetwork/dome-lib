package sql

// Opts are options for requesting of data from the database.
type Opts struct {
	Limit  int
	Offset int
}

// Query used for searching the database.
type Query map[string]string

// SQL provides an interface for compatible SQL based
// database engines such as Postgres, SQLite, MySQL, etc.
type SQL interface {
	Close() error
	Create(interface{}) error
	Delete(interface{}) error
	Exec(string) error
	Get(interface{}) error
	GetWhere(map[string]interface{}, interface{}) error
	Migrate(...interface{}) error
	Open() error
	Select(Query, Opts, interface{}) error
	Update(interface{}) error
}
