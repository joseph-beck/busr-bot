package commands

import (
	"bot/pkg/sql"
	"bot/pkg/util"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func driver(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, invalidTarget, true)
		return
	}

	userId, err := strconv.Atoi(user.ID)
	util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)

	if sql.CheckDriver(userId) {
		driver := sql.Driver(userId)
		msg := fmt.Sprintf(
			">>>  Driver: %s\n University: %s",
			user.Username,
			driver.University,
		)

		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidDriver, true)
	}
}
