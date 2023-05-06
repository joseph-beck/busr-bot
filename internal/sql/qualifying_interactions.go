package sql

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
)

func QualifyingResult(i int, d int) racing.QualifyingResult {
	resultPrimitive := QualifyingResultPrimitive(i, d)

	driver := Driver(resultPrimitive.Driver)
	time := racing.Time{
		Minutes:      resultPrimitive.Minutes,
		Seconds:      resultPrimitive.Seconds,
		Milliseconds: resultPrimitive.Milliseconds,
	}

	return racing.QualifyingResult{
		Id:       i,
		Driver:   driver,
		Position: resultPrimitive.Position,
		Time:     time,
		Points:   resultPrimitive.Points,
	}
}

func QualifyingResultPrimitive(id int, driver int) racing.QualifyingResultPrimitive {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from qualifying_result where id=%d and driver=%d;",
		id, driver,
	))

	util.CheckErr(err)
	defer result.Close()

	var results racing.QualifyingResultPrimitive
	result.Next()
	err = result.StructScan(&results)
	util.CheckErrMsg(err, "at getting qualifying_result primitive")

	return results
}

func AddQualifying(race *racing.QualifyingResult) {
	// conn := Connect()
	// insert, err := conn.db.Query(fmt.Sprintf(
	// 	`insert into quali(id, )
	// 	values(%s);`,
	// 	race.SqlStr(),
	// ))

	// util.CheckErr(err)
	// defer insert.Close()
}

func CheckQualifyingResult(id int, driver int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from qualifying_result where id=%d and driver=%d;",
		id, driver),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
