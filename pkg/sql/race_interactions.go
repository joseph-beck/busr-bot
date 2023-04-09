package sql

import (
	"bot/pkg/f1"
	"bot/pkg/racing"
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

func CheckRace(id int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from race where id=%d",
		id)).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}

func RaceResult(i int, d int) racing.RaceResult {
	resultPrimitive := RaceResultPrimitive(i, d)

	driver := Driver(resultPrimitive.Driver)
	time := racing.Time{
		Minutes:      resultPrimitive.Minutes,
		Seconds:      resultPrimitive.Seconds,
		Milliseconds: resultPrimitive.Milliseconds,
	}

	return racing.RaceResult{
		Id:       i,
		Driver:   driver,
		Position: resultPrimitive.Position,
		Time:     time,
		Points:   resultPrimitive.Points,
	}
}

func RaceResultPrimitive(id int, driver int) racing.RaceResultPrimitive {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from race_result where id=%d and driver=%d;",
		id, driver,
	))

	util.CheckErr(err)
	defer result.Close()

	var results racing.RaceResultPrimitive
	result.Next()
	err = result.StructScan(&results)
	util.CheckErrMsg(err, "at getting race_result primitive")

	return results
}

func CheckRaceResult(id int, driver int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from race_result where id=%d and driver=%d;",
		id, driver),
	).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}
