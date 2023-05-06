package commands

import (
	"github.com/bwmarrin/discordgo"
)

// display winner, tracks, etc

// season
func season(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// sc := i.ApplicationCommandData().Options[0].StringValue()
	// yc := i.ApplicationCommandData().Options[1].StringValue()
	// id := f1.GenSeasonId(sc, yc)

	// if sql.CheckSeason(id) {
	// 	season := sql.GetSeason(id)
	// 	msg := season.Out()
	// 	respond(s, i, msg, false)
	// } else {
	// 	respond(s, i, invalidSeason, true)
	// }
}
