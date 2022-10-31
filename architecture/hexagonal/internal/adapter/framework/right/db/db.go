package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) CloseDBConnection() {
	if err := a.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (a Adapter) Save(answer int32, operation string) error {
	fmt.Printf(`%s %d successfully saved in database`, operation, answer)
	return nil
}
