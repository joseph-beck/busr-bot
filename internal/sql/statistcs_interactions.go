package sql

import (
	"bot/pkg/util"
	"fmt"
)

func (d *Conn) GetWins(id int) int {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	result, err := d.Db.Query(fmt.Sprintf(
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

func (d *Conn) GetPodiums(id int) int {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	result, err := d.Db.Query(fmt.Sprintf(
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

func (d *Conn) GetPoints(id int) float32 {
	d.DbMu.Lock()
	defer d.DbMu.Unlock()

	result, err := d.Db.Query(fmt.Sprintf(
		"select points from driver where id=%d;",
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
