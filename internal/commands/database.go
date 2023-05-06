package commands

import (
	"bot/internal/sql"
	"sync"
)

const path = "configs/database.json"

var lock = &sync.Mutex{}

type database struct {
	Conn sql.Conn
}

var instance *database

func getDatabase() *database {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &database{
			Conn: *sql.MakeConn(path),
		}
	}

	return instance
}
