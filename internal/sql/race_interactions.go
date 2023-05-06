package sql

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
)

func RaceResult(i int, d int) racing.RaceResult {
	resultPrimitive := RaceResultPrimitive(i, d)

	driver := Driver(resultPrimitive.Driver)
	time := racing.Time{
		Minutes:      resultPrimitive.Minutes,
		Seconds:      resultPrimitive.Seconds,
		Milliseconds: resultPrimitive.Milliseconds,
	}

	return racing.RaceResult{
		Id:       i,
		Driver:   driver,
		Position: resultPrimitive.Position,
		Time:     time,
		Points:   resultPrimitive.Points,
	}
}

func RaceResultPrimitive(id int, driver int) racing.RaceResultPrimitive {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from race_result where id=%d and driver=%d;",
		id, driver,
	))

	util.CheckErr(err)
	defer result.Close()

	var results racing.RaceResultPrimitive
	result.Next()
	err = result.StructScan(&results)
	util.CheckErrMsg(err, "at getting race_result primitive")

	return results
}

func AddRaceResult(race racing.RaceResult) {
	// conn := Connect()
	// insert, err := conn.db.Query(fmt.Sprintf(
	// 	`insert into race(id, track, laps, winner, season)
	// 	values(%s);`,
	// 	race.SqlStr(),
	// ))

	// util.CheckErr(err)
	// defer insert.Close()
}

func CheckRaceResult(id int, driver int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from race_result where id=%d and driver=%d;",
		id, driver),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
