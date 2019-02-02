package db

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type Connection struct {
	Props *Props
	Conn *sql.DB
}

func (c *Connection) Connect() error {
	if c.Props == nil {
		return errors.New("DB_NO_CONN_PROPS")
	}
	
}

