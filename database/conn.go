package database

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // add this
)

type Connection struct {
	db *sqlx.DB
}

var ErrOpenConnection = errors.New("problem opening db connection")
var ErrCreate = errors.New("problem inserting into table")
var ErrRetrieving = errors.New("problem retrieving from table")

func NewConnection() (*Connection, error) {
	connStr := "postgresql://postgres:example@localhost/school?sslmode=disable"

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrOpenConnection, err)
	}

	conn := &Connection{
		db: db,
	}

	return conn, nil
}
