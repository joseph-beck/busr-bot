package sql

import (
	"bot/pkg/util"
	"database/sql"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type conn struct {
	dbMu sync.Mutex
	db   sqlx.DB
}

func MakeDbx(path string) *conn {
	return &conn{
		db: ConnDbx(path),
	}
}

func (d *conn) Close() {
	d.dbMu.Lock()
	defer d.dbMu.Unlock()

	defer d.db.Close()
}

func ConnDbx(path string) sqlx.DB {
	db, err := sqlx.Open("mysql", connStr(path))
	util.CheckErrMsg(err, "SQL connection failure")
	log.Println("Database connected")

	return *db
}

func ConnDb(path string) sql.DB {
	db, err := sql.Open("mysql", connStr(path))
	util.CheckErrMsg(err, "SQL connection failure")
	log.Println("Database connected")

	return *db
}
