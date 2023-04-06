package main

import (
	"bot/pkg/sql"
	"bot/pkg/f1"
	"fmt"
)

func main() {
	fmt.Println("running tester")

	driver := sql.GetDriver(286245298010062851)
	fmt.Println(driver.Output())

	var dasher = f1.Driver{
		Id: 313014742753214465,
		Name: "dasher",
		University: "Reading",
		Wins: 2,
		Poles: 5,
		Podiums: 12,
		Starts: 65,
		Points: 5681.75,
		Avg_quali: 4.35,
	}

	sql.AddDriver(dasher)

	defer sql.Disconnect()
}
