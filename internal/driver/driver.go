package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 5
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	err = testDb(d)
	if err != nil {
		return nil, err
	}
	
	dbConn.SQL = d
	return dbConn, nil
}

func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	fmt.Println("connected to db")
	
	return nil
}


