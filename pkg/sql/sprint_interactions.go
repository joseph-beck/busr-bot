package sql

import (
	"bot/pkg/racing"
	"bot/pkg/util"
	"fmt"
)

func SprintResult(i int, d int) racing.SprintResult {
	resultPrimitive := SprintResultPrimitive(i, d)

	driver := Driver(resultPrimitive.Driver)
	time := racing.Time{
		Minutes:      resultPrimitive.Minutes,
		Seconds:      resultPrimitive.Seconds,
		Milliseconds: resultPrimitive.Milliseconds,
	}

	return racing.SprintResult{
		Id:       i,
		Driver:   driver,
		Position: resultPrimitive.Position,
		Time:     time,
		Points:   resultPrimitive.Points,
	}
}

func SprintResultPrimitive(id int, driver int) racing.SprintResultPrimitive {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from sprint_result where id=%d and driver=%d;",
		id, driver,
	))

	util.CheckErr(err)
	defer result.Close()

	var results racing.SprintResultPrimitive
	result.Next()
	err = result.StructScan(&results)
	util.CheckErrMsg(err, "at getting sprint_result primitive")

	return results
}

func CheckSprintResult(id int, driver int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from sprint_result where id=%d and driver=%d;",
		id, driver),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
