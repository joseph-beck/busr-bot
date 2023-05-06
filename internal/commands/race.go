package commands

import (
	"bot/internal/racing"
	"bot/internal/sql"
	"bot/pkg/util"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func race(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, invalidTarget, true)
		return
	}

	series := i.ApplicationCommandData().Options[1].StringValue()
	race := i.ApplicationCommandData().Options[2].StringValue()
	season := i.ApplicationCommandData().Options[3].StringValue()
	year := i.ApplicationCommandData().Options[4].StringValue()

	userId, err := strconv.Atoi(user.ID)
	util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)
	raceResultId := racing.GenRaceId(series, race, season, year)

	if sql.CheckRaceResult(raceResultId, userId) {
		raceResult := sql.RaceResult(raceResultId, userId)
		msg := fmt.Sprintf(
			">>> **Race result for %s**: \n%s",
			user.Username,
			raceResult.Out(),
		)
		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidRace, true)
	}
}
