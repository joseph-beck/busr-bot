package f1

import (
	"bot/pkg/util"
	"strconv"
)

func GenSeasonId(season string, year string) int {
	i := "5"
	i += season + year

	id, err := strconv.Atoi(i)
	util.CheckErrMsg(err, "Error generating season id")
	return id
}

func GenRaceId(race string, year string, season string) int {
	i := "4"
	i += race + year + season

	id, err := strconv.Atoi(i)
	util.CheckErrMsg(err, "Error generating race id")
	return id
}

func GenQualifyingId(race string, year string, season string) int {
	i := "9"
	i += race + year + season + "1"

	id, err := strconv.Atoi(i)
	util.CheckErrMsg(err, "Error generating qualifying id")
	return id
}
