package commands

import (
	"bot/pkg/racing"
	"bot/pkg/sql"

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
	id := racing.GenWeekendId(series, race, season, year)
	
	if sql.CheckWeekend(id) {
		msg := sql.WeekendPrimitive(id).Str()
		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidRace, true)
	}
}

