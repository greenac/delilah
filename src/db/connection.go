package db

import (
	"database/sql"
	"errors"
	"github.com/greenac/artemis/logger"
	_ "github.com/lib/pq"
)

type Type string
const (
	Postgres Type = "postgres"
)

type Connection struct {
	Props *Props
	Db *sql.DB
}

func (c *Connection) Connect() error {
	if c.Props == nil {
		return errors.New("DB_NO_CONN_PROPS")
	}

	db, err := sql.Open(string(Postgres), c.Props.ConnectionString()); if err != nil {
		logger.Error("`Connection::Connect` creating connection:", err)
		return err
	}

	c.Db = db

	return nil
}

func (c *Connection) FindById(id int, table string) (*sql.Rows, error) {
	rows, err := c.Db.Query("SELECT * FROM $1 WHERE id=$2", table, id); if err != nil {
		logger.Error("`Connection:FindById` making query:", err)
		return nil, err
	}

	return rows, nil
}

