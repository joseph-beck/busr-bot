package main

import (
	discord "bot/pkg/disgo"
	"bot/pkg/sql"
	"log"
)

func main() {
	log.Println("Running tester")
	sql.Connect()
	discord.OpenSession()
	defer sql.Disconnect()
}
