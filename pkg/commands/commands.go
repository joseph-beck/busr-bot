package commands

import (
	"bot/pkg/util"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var cmds = []*discordgo.ApplicationCommand{
	{
		Name:        "stats",
		Description: "Get another players stats.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "Player who's stats you wish to retrieve.",
				Required:    true,
			},
		},
	},
	{
		Name:        "wins",
		Description: "Get another players wins.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "Player who's wins you wish to retrieve.",
				Required:    true,
			},
		},
	},
	{
		Name:        "podiums",
		Description: "Get another players podiums.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "Player who's podiums you wish to retrieve.",
				Required:    true,
			},
		},
	},
	{
		Name:        "points",
		Description: "Get another players points.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "Player who's points you wish to retrieve.",
				Required:    true,
			},
		},
	},
	{
		Name:        "season",
		Description: "Get a season result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "season",
				Description: "Season (Spring, Summer or Winter)",
				Required:    true,
				Choices:     seasonChoices,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "year",
				Description: "Year of Season",
				Required:    true,
				Choices:     yearChoices,
			},
		},
	},
	{
		Name:        "race",
		Description: "Get a race result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "race",
				Description: "Name of track.",
				Required:    true,
				Choices:     raceChoices,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "season",
				Description: "Season of race.",
				Required:    true,
				Choices:     seasonChoices,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "year",
				Description: "Year of race.",
				Required:    true,
				Choices:     yearChoices,
			},
		},
	},
	{
		Name:        "quali",
		Description: "Get a qualifying result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "race",
				Description: "Name of track.",
				Required:    true,
				Choices:     raceChoices,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "season",
				Description: "Season of race.",
				Required:    true,
				Choices:     seasonChoices,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "year",
				Description: "Year of race.",
				Required:    true,
				Choices:     yearChoices,
			},
		},
	},
}

var regCmds []*discordgo.ApplicationCommand

var cmdHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"stats":   stats,
	"wins":    wins,
	"podiums": podiums,
	"points":  points,
	"season":  season,
	"race":    race,
	"quali":   qualifying,
}

var btnHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

func RegisterCommands(s *discordgo.Session) {
	regCmds = make([]*discordgo.ApplicationCommand, len(cmds))

	for i, cmd := range cmds {
		c, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		util.CheckErrMsg(err, "Command creation failure :/"+cmd.Name)

		regCmds[i] = c
		log.Printf(" - /%s", cmd.Name)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
