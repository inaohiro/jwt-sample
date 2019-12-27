package persistence

import (
	"database/sql"
	"log"
	"time"

	"github.com/inaohiro/jwt-sample/config"
)

type DB struct {
	*sql.DB
}

func NewConnection(c *config.DBEnv) (*DB, func(), error) {
	db, err := sql.Open(c.Driver, c.DataSource())
	if err != nil {
		return nil, nil, err
	}
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}
	return &DB{db}, cleanup, nil
}
