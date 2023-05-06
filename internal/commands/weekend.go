package commands

import (
	"bot/internal/racing"
	"bot/internal/sql"
	"bot/pkg/util"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// weekend
func weekend(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	weekendId := racing.GenWeekendId(series, race, season, year)

	if sql.CheckWeekend(weekendId) {
		weekend := sql.WeekendPrimitive(weekendId)
		msg := fmt.Sprintf(
			">>> **%s's race weekend**\n",
			user.Username,
		)

		if sql.CheckQualifyingResult(weekend.Qualifying, userId) {
			qualifyingResult := sql.QualifyingResult(weekend.Qualifying, userId)
			msg += "\n **Qualifying**: " + qualifyingResult.Out()
		}

		if sql.CheckSprintResult(weekend.Sprint, userId) {
			sprintResult := sql.SprintResult(weekend.Sprint, userId)
			msg += "\n **Sprint**: " + sprintResult.Out()
		}

		if sql.CheckRaceResult(weekend.Race, userId) {
			raceResult := sql.RaceResult(weekend.Race, userId)
			msg += "\n **Race**: " + raceResult.Out()
		}

		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidRace, true)
	}
}
