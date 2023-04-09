package commands

import (
	"bot/pkg/racing"
	"bot/pkg/sql"
	"bot/pkg/util"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func sprint(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	raceResultId := racing.GenSprintId(series, race, season, year)

	if sql.CheckSprintResult(raceResultId, userId) {
		raceResult := sql.SprintResult(raceResultId, userId)
		msg := fmt.Sprintf(
			">>> **Sprint result for %s**: \n%s",
			user.Username,
			raceResult.Out(),
		)
		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidSprint, true)
	}
}
