package commands

import (
	"bot/internal/racing"
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

	c := getDatabase()
	if c.Conn.CheckWeekend(weekendId) {
		weekend := c.Conn.WeekendPrimitive(weekendId)
		msg := fmt.Sprintf(
			">>> **%s's race weekend**\n",
			user.Username,
		)

		if c.Conn.CheckQualifyingResult(weekend.Qualifying, userId) {
			qualifyingResult := c.Conn.QualifyingResult(weekend.Qualifying, userId)
			msg += "\n **Qualifying**: " + qualifyingResult.Out()
		}

		if c.Conn.CheckSprintResult(weekend.Sprint, userId) {
			sprintResult := c.Conn.SprintResult(weekend.Sprint, userId)
			msg += "\n **Sprint**: " + sprintResult.Out()
		}

		if c.Conn.CheckRaceResult(weekend.Race, userId) {
			raceResult := c.Conn.RaceResult(weekend.Race, userId)
			msg += "\n **Race**: " + raceResult.Out()
		}

		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidRace, true)
	}
}
