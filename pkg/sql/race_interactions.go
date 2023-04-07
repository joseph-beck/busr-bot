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
	var podium [3]int
	var season int

	result.Next()
	err = result.Scan(
		&race.Id,
		&race.Track,
		&race.Laps,
		&podium[0],
		&podium[1],
		&podium[2],
		&season,
	)
	util.CheckErr(err)

	race.First = GetDriver(podium[0])
	race.Second = GetDriver(podium[1])
	race.Third = GetDriver(podium[2])
	race.Season = GetSeason(season)

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
