package sql

import (
	"bot/pkg/f1"
	"bot/pkg/util"
	"fmt"
)

// add stuff here i guess

func GetDriver(id int) f1.Driver {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from driver where id=%d;",
		id))

	util.CheckErr(err)
	defer result.Close()

	var driver f1.Driver
	result.Next()
	err = result.StructScan(&driver)
	util.CheckErr(err)

	return driver
}

func AddDriver(driver f1.Driver) {
	conn := Connect()
	insert, err := conn.db.Query(fmt.Sprintf(
		`insert into driver(id, name, university, wins, poles, podiums, starts, points, avg_quali) 
		values(%s);`,
		driver.SqlStr()))

	util.CheckErr(err)
	defer insert.Close()
}

func UpdateDriver(driver *f1.Driver) {
	conn := Connect()
	updates := driver.UpdateStr()

	query := fmt.Sprintf(
		"id='%d'",
		driver.Id)

	update, err := conn.db.Query(fmt.Sprintf(
		`update driver
		set %s
		where %s;`,
		updates,
		query))

	util.CheckErr(err)
	defer update.Close()
}