package commands

import (
	"bot/pkg/sql"

	"github.com/bwmarrin/discordgo"
)

func race(s *discordgo.Session, i *discordgo.InteractionCreate) {
	race := sql.GetRace(4142301)
	msg := race.Out()
	respond(s, i, msg, false)
}
