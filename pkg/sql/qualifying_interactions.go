package sql

import (
	"bot/pkg/f1"
	"bot/pkg/util"
	"fmt"
)

func GetQualifying(id int) f1.Qualifying {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from qualifying where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var quali f1.Qualifying
	var resultId int
	var timeId int

	result.Next()
	err = result.Scan(
		&quali.Id,
		&quali.Track,
		&resultId,
		&timeId,
	)
	util.CheckErrMsg(err, "at getting qualifying")

	quali.Results = GetQualifyingResults(resultId)
	quali.Timings = GetQualifyingTimings(timeId)

	return quali
}

func GetQualifyingResults(id int) f1.QualifyingResults {
	fmt.Println(id)
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from qualifying_results where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var results f1.QualifyingResults
	var podium [3]int

	result.Next()
	err = result.Scan(
		&results.Id,
		&podium[0],
		&podium[1],
		&podium[2],
	)
	util.CheckErrMsg(err, "at getting qualifying results")

	results.First = GetDriver(podium[0])
	results.Second = GetDriver(podium[1])
	results.Third = GetDriver(podium[2])

	return results
}

func GetQualifyingTimings(id int) f1.QualifyingTimings {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from qualifying_results where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var timings f1.QualifyingTimings

	result.Next()
	err = result.Scan(
		&timings.Id,
		&timings.First.Minutes,
		&timings.First.Seconds,
		&timings.First.Miliseconds,
		&timings.Second.Minutes,
		&timings.Second.Seconds,
		&timings.Second.Miliseconds,
		&timings.Third.Minutes,
		&timings.Third.Seconds,
		&timings.Third.Miliseconds,
	)
	util.CheckErrMsg(err, "at getting qualifying results")

	return timings
}

func AddQualifying(race f1.Race) {
	conn := Connect()
	insert, err := conn.db.Query(fmt.Sprintf(
		`insert into quali(id, ) 
		values(%s);`,
		race.SqlStr(),
	))

	util.CheckErr(err)
	defer insert.Close()
}

func CheckQualifying(id int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from qualifying where id=%d",
		id)).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}