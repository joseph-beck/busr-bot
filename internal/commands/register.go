package commands

import (
	"bot/internal/sql"
	"bot/pkg/util"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session) {
	regCmds = make([]*discordgo.ApplicationCommand, len(cmds))

	for i, cmd := range cmds {
		c, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		util.CheckErrMsg(err, "Command creation failure :/"+cmd.Name)

		regCmds[i] = c
		log.Printf(" - /%s", cmd.Name)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate, db *sql.Conn) {
		if i.Type == discordgo.InteractionApplicationCommand {
			h, exists := cmdHandlers[i.ApplicationCommandData().Name]
			if exists {
				h(s, i)
			}
		} else if i.Type == discordgo.InteractionMessageComponent {
			h, exists := btnHandlers[i.MessageComponentData().CustomID]
			if exists {
				h(s, i)
			}
		}
	})
}

func ClearCommands(s *discordgo.Session) {
	for _, cmd := range regCmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		util.CheckErrMsg(err, fmt.Sprintf("Delete command failure /%s", cmd.Name))
		log.Printf(" - /%s", cmd.Name)
	}
}

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, message string, hidden bool) {
	var flag discordgo.MessageFlags

	if hidden {
		flag = discordgo.MessageFlagsEphemeral
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   flag,
		},
	})
}
