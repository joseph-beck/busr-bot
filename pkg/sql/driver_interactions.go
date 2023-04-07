package sql

import (
	"bot/pkg/f1"
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
		"select id from driver where id=%d",
		id)).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}

func GetWins(id int) int {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select wins from driver where id=%d",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var wins int
	result.Next()
	err = result.Scan(&wins)
	util.CheckErr(err)

	return wins
}

func GetPodiums(id int) int {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select podiums from driver where id=%d",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var podiums int
	result.Next()
	err = result.Scan(&podiums)
	util.CheckErr(err)

	return podiums
}

func GetPoints(id int) float32 {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select points from driver where id=%d",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var points float32
	result.Next()
	err = result.Scan(&points)
	util.CheckErr(err)

	return points
}
