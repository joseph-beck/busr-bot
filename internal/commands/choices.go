package commands

import "github.com/bwmarrin/discordgo"

var (
	formulaChoices = []*discordgo.ApplicationCommandOptionChoice{
		{
			Name:  "Formula 1",
			Value: "1",
		},
		{
			Name:  "Formula 2",
			Value: "2",
		},
	}

	seasonChoices = []*discordgo.ApplicationCommandOptionChoice{
		{
			Name:  "Spring",
			Value: "01",
		},
		{
			Name:  "Summer",
			Value: "02",
		},
		{
			Name:  "Winter",
			Value: "03",
		},
	}

	yearChoices = []*discordgo.ApplicationCommandOptionChoice{
		{
			Name:  "22",
			Value: "22",
		},
		{
			Name:  "23",
			Value: "23",
		},
	}

	raceChoices = []*discordgo.ApplicationCommandOptionChoice{
		{
			Name:  "Bahrain",
			Value: "01",
		},
		{
			Name:  "Saudi Arabia",
			Value: "02",
		},
		{
			Name:  "Australia",
			Value: "03",
		},
		{
			Name:  "Imola",
			Value: "04",
		},
		{
			Name:  "Miami",
			Value: "05",
		},
		{
			Name:  "Spain",
			Value: "06",
		},
		{
			Name:  "Monaco",
			Value: "07",
		},
		{
			Name:  "Azerbaijan",
			Value: "08",
		},
		{
			Name:  "Canada",
			Value: "09",
		},
		{
			Name:  "Great Britain",
			Value: "10",
		},
		{
			Name:  "Austria",
			Value: "11",
		},
		{
			Name:  "France",
			Value: "12",
		},
		{
			Name:  "Hungary",
			Value: "13",
		},
		{
			Name:  "Belgium",
			Value: "14",
		},
		{
			Name:  "Netherlands",
			Value: "15",
		},
		{
			Name:  "Monza",
			Value: "16",
		},
		{
			Name:  "Singapore",
			Value: "17",
		},
		{
			Name:  "Japan",
			Value: "18",
		},
		{
			Name:  "Austin",
			Value: "19",
		},
		{
			Name:  "Mexico",
			Value: "20",
		},
		{
			Name:  "Brazil",
			Value: "21",
		},
		{
			Name:  "Abu Dhabi",
			Value: "22",
		},
		{
			Name:  "Portugal",
			Value: "23",
		},
		{
			Name:  "China",
			Value: "24",
		},
	}
)
