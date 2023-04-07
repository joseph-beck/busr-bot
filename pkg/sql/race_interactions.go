package sql

import (
	"bot/pkg/f1"
	"bot/pkg/util"
	"fmt"
)

func GetRace(id int) f1.Race {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from race where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var race f1.Race

	var pIds [3]int
	var sId int

	result.Next()
	err = result.Scan(
		&race.Id,
		&race.Track,
		&race.Laps,
		&pIds[0],
		&pIds[1],
		&pIds[2],
		&sId,
	)
	util.CheckErr(err)

	race.First = GetDriver(pIds[0])
	race.Second = GetDriver(pIds[1])
	race.Third = GetDriver(pIds[2])
	race.Season = GetSeason(sId)

	return race
}

func AddRace(race f1.Race) {
	conn := Connect()
	insert, err := conn.db.Query(fmt.Sprintf(
		`insert into race(id, track, laps, winner, season) 
		values(%s);`,
		race.SqlStr(),
	))

	util.CheckErr(err)
	defer insert.Close()
}
