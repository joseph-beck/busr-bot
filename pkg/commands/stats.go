package commands

import (
	"bot/pkg/sql"
	"bot/pkg/util"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// driver

func stats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if !user.Bot && user.ID != i.Interaction.Member.User.ID {

		id, err := strconv.Atoi(user.ID)
		util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)

		if sql.CheckDriver(id) {
			driver := sql.GetDriver(id)

			m := driver.Out()

			respond(s, i, m, false)
		} else {
			respond(s, i, "Driver does not exist!", true)
		}
	} else {
		respond(s, i, "Not a valid target", true)
	}
}

// wins

// podiums

// points
