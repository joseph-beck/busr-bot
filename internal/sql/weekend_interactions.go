package sql

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
)

func WeekendPrimitive(id int) racing.RaceWeekendPrimitive {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from race_weekend where id=%d;",
		id,
	))
	util.CheckErr(err)
	defer result.Close()

	var weekend racing.RaceWeekendPrimitive
	result.Next()
	err = result.StructScan(&weekend)
	util.CheckErr(err)

	return weekend
}

func CheckWeekend(id int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from race_weekend where id=%d;",
		id),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
