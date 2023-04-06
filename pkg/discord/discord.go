package discord

import (
	"bot/pkg/util"
	"log"

	"github.com/bwmarrin/discordgo"
)

func OpenSession() {
	conf := Config()
	session, err := discordgo.New("Bot" + conf.Token)
	util.CheckErrMsg(err, "Discordgo bot init failure")

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in %s : %s", r.User.Username, r.User.Discriminator)
	})

	err = session.Open()
	util.CheckErrMsg(err, "Session opening failure")
	defer session.Close()
}
