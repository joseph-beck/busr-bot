package commands

import (
	"bot/pkg/f1"
	"bot/pkg/sql"

	"github.com/bwmarrin/discordgo"
)

// qualifying
func qualifying(s *discordgo.Session, i *discordgo.InteractionCreate) {
	qc := i.ApplicationCommandData().Options[0].StringValue()
	sc := i.ApplicationCommandData().Options[1].StringValue()
	yc := i.ApplicationCommandData().Options[2].StringValue()
	id := f1.GenQualifyingId(qc, yc, sc)

	if sql.CheckQualifying(id) {
		quali := sql.GetQualifying(id)
		msg := quali.Out()
		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidRace, true)
	}
}
