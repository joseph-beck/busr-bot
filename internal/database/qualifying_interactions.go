package database

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
)

func (d *Conn) QualifyingResult(id int, drvr int) racing.QualifyingResult {
	resultPrimitive := d.QualifyingResultPrimitive(id, drvr)

	driver := d.Driver(resultPrimitive.Driver)
	time := racing.Time{
		Minutes:      resultPrimitive.Minutes,
		Seconds:      resultPrimitive.Seconds,
		Milliseconds: resultPrimitive.Milliseconds,
	}

	return racing.QualifyingResult{
		Id:       id,
		Driver:   driver,
		Position: resultPrimitive.Position,
		Time:     time,
		Points:   resultPrimitive.Points,
	}
}

func (d *Conn) QualifyingResultPrimitive(id int, driver int) racing.QualifyingResultPrimitive {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	result, err := d.Db.Queryx(fmt.Sprintf(
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

func (d *Conn) AddQualifying(race *racing.QualifyingResult) {
	// conn := Connect()
	// insert, err := conn.db.Query(fmt.Sprintf(
	// 	`insert into quali(id, )
	// 	values(%s);`,
	// 	race.SqlStr(),
	// ))

	// util.CheckErr(err)
	// defer insert.Close()
}

func (d *Conn) CheckQualifyingResult(id int, driver int) bool {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	err := d.Db.QueryRow(fmt.Sprintf(
		"select id from qualifying_result where id=%d and driver=%d;",
		id, driver),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
