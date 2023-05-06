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

func MakeDb() *conn {
	return &conn{
		db: ConnDbx(),
	}
}

func (d *conn) Close() {
	d.dbMu.Lock()
	defer d.dbMu.Unlock()

	defer d.db.Close()
}

func ConnDbx() sqlx.DB {
	db, err := sqlx.Open("mysql", connStr())
	util.CheckErrMsg(err, "SQL connection failure")
	log.Println("Database connected")

	return *db
}

func ConnDb() sql.DB {
	db, err := sql.Open("mysql", connStr())
	util.CheckErrMsg(err, "SQL connection failure")
	log.Println("Database connected")

	return *db
}
