package discord

import (
	"bot/internal/commands"
	"bot/pkg/util"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func OpenSession() {
	conf := Config()
	session, err := discordgo.New("Bot " + conf.Token)
	util.CheckErrMsg(err, "Discordgo bot init failure")

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s#%s", r.User.Username, r.User.Discriminator)
	})

	err = session.Open()
	util.CheckErrMsg(err, "Session opening failure")
	defer session.Close()

	log.Println("Registering slash cmds")
	commands.RegisterCommands(session)
	log.Println("Commands registered")

	err = session.UpdateGameStatus(0, conf.Status)
	util.CheckErrMsg(err, "Status update failed, status: "+conf.Status)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Clearing slash cmds")
	commands.ClearCommands(session)
	log.Println("Commands cleared")
}
