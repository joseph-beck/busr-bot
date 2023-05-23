package database

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
)

func (d *Conn) Driver(id int) racing.Driver {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	result, err := d.Db.Queryx(fmt.Sprintf(
		"select * from drivers where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var driver racing.Driver
	result.Next()
	err = result.StructScan(&driver)
	util.CheckErr(err)

	return driver
}

func (d *Conn) AddDriver(driver *racing.Driver) {
	// conn := Connect()
	// insert, err := conn.db.Query(fmt.Sprintf(
	// 	`insert into driver(id, name, university, wins, poles, podiums, starts, points, avg_quali)
	// 	values(%s);`,
	// 	driver.SqlStr(),
	// ))

	// util.CheckErr(err)
	// defer insert.Close()
}

func (d *Conn) UpdateDriver(driver *racing.Driver) {
	// conn := Connect()
	// updates := driver.UpdateStr()

	// query := fmt.Sprintf(
	// 	"id='%d'",
	// 	driver.Id)

	// update, err := conn.db.Query(fmt.Sprintf(
	// 	`update drivers
	// 	set %s
	// 	where %s;`,
	// 	updates,
	// 	query,
	// ))

	// util.CheckErr(err)
	// defer update.Close()
}

func (d *Conn) CheckDriver(id int) bool {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	err := d.Db.QueryRow(fmt.Sprintf(
		"select id from drivers where id=%d;",
		id),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
