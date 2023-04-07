package commands

import (
	"bot/pkg/sql"

	"github.com/bwmarrin/discordgo"
)

// display winner, tracks, etc

// season
func season(s *discordgo.Session, i *discordgo.InteractionCreate) {
	season := sql.GetSeason(50123)
	msg := season.Out()
	respond(s, i, msg, false)
}
