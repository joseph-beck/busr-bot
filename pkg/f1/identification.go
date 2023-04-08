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
	util.CheckErrMsg(err, "Error generating season id")
	return id
}
