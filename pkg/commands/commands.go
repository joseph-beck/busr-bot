package commands

import "github.com/bwmarrin/discordgo"

var cmds = []*discordgo.ApplicationCommand{
	{
		Name:        "driver",
		Description: "Get a driver",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "What driver do you want to see?",
				Required:    true,
			},
		},
	},
	{
		Name:        "quali",
		Description: "Get a qualifying result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "What driver do you want to see?",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "formula",
				Description: "What formula series was the race?",
				Required:    true,
				Choices:     formulaChoices,
			},
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
		Name:        "sprint",
		Description: "Get a sprint result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "What driver do you want to see?",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "formula",
				Description: "What formula series was the race?",
				Required:    true,
				Choices:     formulaChoices,
			},
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
		Name:        "race",
		Description: "Get a race result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "What driver do you want to see?",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "formula",
				Description: "What formula series was the race?",
				Required:    true,
				Choices:     formulaChoices,
			},
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
		Name:        "weekend",
		Description: "Get a weekend result.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "driver",
				Description: "What driver do you want to see?",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "formula",
				Description: "What formula series was the race?",
				Required:    true,
				Choices:     formulaChoices,
			},
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
