package sql

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
)

func (d *Conn) WeekendPrimitive(id int) racing.RaceWeekendPrimitive {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	result, err := d.Db.Queryx(fmt.Sprintf(
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

func (d *Conn) CheckWeekend(id int) bool {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	err := d.Db.QueryRow(fmt.Sprintf(
		"select id from race_weekend where id=%d;",
		id),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
