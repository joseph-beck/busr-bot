package sql

import (
	"bot/pkg/f1"
	"fmt"
)

// add stuff here i guess

func GetDriver(id int) f1.Driver {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from driver where id=%d;",
		id))

	checkErr(err)
	defer result.Close()

	var driver f1.Driver
	result.Next()
	err = result.StructScan(&driver)
	checkErr(err)

	return driver
}

func AddDriver(driver f1.Driver) {
	conn := Connect()
	insert, err := conn.db.Query(fmt.Sprintf(
		`insert into driver(id, name, university, wins, poles, podiums, starts, points, avg_quali) 
		values(%s);`,
		driver.String()))

	checkErr(err)
	defer insert.Close()
}
