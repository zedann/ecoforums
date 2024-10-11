package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB // private prop
}

type databaseConfig struct {
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
}

func New() (*Database, error) {

	dbConf := databaseConfig{
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabaseUser:     os.Getenv("DB_USER"),
		DatabasePassword: os.Getenv("DB_PASS"),
		DatabaseHost:     os.Getenv("DB_HOST"),
	}

	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", dbConf.DatabaseUser, dbConf.DatabasePassword, dbConf.DatabaseHost, dbConf.DatabaseName)
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil

}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
