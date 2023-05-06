package sql

import (
	s "bot/pkg/sql"
	"sync"

	"github.com/jmoiron/sqlx"
)

type Conn struct {
	DbMu sync.Mutex
	Db   sqlx.DB
}

func MakeConn(path string) *Conn {
	return &Conn{
		Db: s.ConnDbx(path),
	}
}

func (d *Conn) Close() {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	defer d.Db.Close()
}
