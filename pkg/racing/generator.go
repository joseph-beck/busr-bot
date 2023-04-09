package racing

import (
	"bot/pkg/util"
	"strconv"
)

func GenWeekendId(series string, race string, season string, year string) int {
	i := series + "1" + race + season + year
	id, err := strconv.Atoi(i)
	util.CheckErrMsg(err, "Error generating weekend id")
	return id
}

func GenRaceId(series string, race string, season string, year string) int {
	i := series + "4" + race + season + year
	id, err := strconv.Atoi(i)
	util.CheckErrMsg(err, "Error generating race id")
	return id
}

func GenQualiId(series string, race string, season string, year string) int {
	i := series + "9" + race + season + year
	id, err := strconv.Atoi(i)
	util.CheckErrMsg(err, "Error generating qualifying id")
	return id
}