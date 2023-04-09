package main

import (
	"bot/pkg/discord"
	"bot/pkg/sql"
	"log"
)

func main() {
	log.Println("Running tester")
	sql.Connect()
	discord.OpenSession()
	//defer sql.Disconnect()
}
