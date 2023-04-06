package main

import (
	"bot/pkg/sql"
	"fmt"
)

func main() {
	fmt.Println("running tester")

	maddog := sql.GetDriver(286245298010062851)
	fmt.Println(maddog.Out())

	maddog.University = "Leicester"

	sql.UpdateDriver(&maddog)

	defer sql.Disconnect()
}
