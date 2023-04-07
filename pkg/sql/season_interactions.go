package sql

import (
	"bot/pkg/f1"
	"bot/pkg/util"
	"fmt"
)

func GetSeason(id int) f1.Season {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from season where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var season f1.Season
	var champion int

	result.Next()
	err = result.Scan(
		&season.Id,
		&season.Year,
		&season.Season,
		&champion,
	)
	util.CheckErr(err)

	season.Champion = GetDriver(champion)

	return season
}

func AddSeason(race f1.Race) {
	conn := Connect()
	insert, err := conn.db.Query(fmt.Sprintf(
		`insert into season(id, name, university, wins, poles, podiums, starts, points, avg_quali) 
		values(%s);`,
		race.SqlStr(),
	))

	util.CheckErr(err)
	defer insert.Close()
}
