package sql

import (
	d "bot/pkg/database"
	"sync"

	"github.com/jmoiron/sqlx"
)

type Conn struct {
	DbMu sync.Mutex
	Db   sqlx.DB
}

func MakeConn(path string) *Conn {
	return &Conn{
		Db: d.ConnDbx(path),
	}
}

func (d *Conn) Close() {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	defer d.Db.Close()
}
