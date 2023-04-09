package sql

import (
	"bot/pkg/f1"
	"bot/pkg/racing"
	"bot/pkg/util"
	"fmt"
)

func GetDriver(id int) f1.Driver {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from driver where id=%d;",
		id,
	))

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
		driver.SqlStr(),
	))

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
		query,
	))

	util.CheckErr(err)
	defer update.Close()
}

func CheckDriver(id int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from drivers where id=%d;",
		id),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}


func Driver(id int) racing.Driver {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
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
