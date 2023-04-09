package sql

import (
	"bot/pkg/util"
	"fmt"
)

func GetWins(id int) int {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select wins from driver where id=%d",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var wins int
	result.Next()
	err = result.Scan(&wins)
	util.CheckErr(err)

	return wins
}

func GetPodiums(id int) int {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select podiums from driver where id=%d",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var podiums int
	result.Next()
	err = result.Scan(&podiums)
	util.CheckErr(err)

	return podiums
}

func GetPoints(id int) float32 {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select points from driver where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var points float32
	result.Next()
	err = result.Scan(&points)
	util.CheckErr(err)

	return points
}
