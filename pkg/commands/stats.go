package commands

import (
	"fmt"
	"time"

	"bot/pkg/discord"
	"bot/pkg/f1"

	"github.com/bwmarrin/discordgo"
)

// driver

func stats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if !user.Bot && user.ID != i.Interaction.Member.User.ID {

	}
}

// wins

// podiums

// points