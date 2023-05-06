package commands

import "github.com/bwmarrin/discordgo"

var cmdHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"driver":  driver,
	"quali":   qualifying,
	"sprint":  sprint,
	"race":    race,
	"weekend": weekend,
}

var btnHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
