package commands

import (
	"bot/internal/database"
	"sync"
)

const path = "configs/database.json"

var lock = &sync.Mutex{}

type db struct {
	Conn database.Conn
}

var instance *db

func getDatabase() *db {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &db{
			Conn: *database.MakeConn(path),
		}
	}

	return instance
}
