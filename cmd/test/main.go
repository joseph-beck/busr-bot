package main

import (
	"bot/pkg/discord"
	"bot/pkg/sql"
	"fmt"
)

func main() {
	fmt.Println("running tester")

	maddog := sql.GetDriver(286245298010062851)
	fmt.Println(maddog.Out())

	defer sql.Disconnect()

	discord.OpenSession()
}
