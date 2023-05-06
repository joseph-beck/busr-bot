package sql

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
)

func (d *Conn) SprintResult(id int, drvr int) racing.SprintResult {
	resultPrimitive := d.SprintResultPrimitive(id, drvr)

	driver := d.Driver(resultPrimitive.Driver)
	time := racing.Time{
		Minutes:      resultPrimitive.Minutes,
		Seconds:      resultPrimitive.Seconds,
		Milliseconds: resultPrimitive.Milliseconds,
	}

	return racing.SprintResult{
		Id:       id,
		Driver:   driver,
		Position: resultPrimitive.Position,
		Time:     time,
		Points:   resultPrimitive.Points,
	}
}

func (d *Conn) SprintResultPrimitive(id int, driver int) racing.SprintResultPrimitive {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	result, err := d.Db.Queryx(fmt.Sprintf(
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

func (d *Conn) CheckSprintResult(id int, driver int) bool {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	err := d.Db.QueryRow(fmt.Sprintf(
		"select id from sprint_result where id=%d and driver=%d;",
		id, driver),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
