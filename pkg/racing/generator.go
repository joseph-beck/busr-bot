package racing

import (
	"bot/pkg/util"
	"strconv"
)

func GenWeekendId(series string, race string, season string, year string) int {
	i := series + "1" + race + season + year
	id, err := strconv.Atoi(i)
	util.CheckErrMsg(err, "Error generating race id")
	return id
}
