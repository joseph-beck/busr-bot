package commands

import (
	"bot/pkg/f1"
	"bot/pkg/sql"

	"github.com/bwmarrin/discordgo"
)

func race(s *discordgo.Session, i *discordgo.InteractionCreate) {
	rc := i.ApplicationCommandData().Options[0].StringValue()
	sc := i.ApplicationCommandData().Options[1].StringValue()
	yc := i.ApplicationCommandData().Options[2].StringValue()
	id := f1.GenRaceId(rc, yc, sc)

	if sql.CheckRace(id) {
		race := sql.GetRace(id)
		msg := race.Out()
		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidRace, true)
	}
}
