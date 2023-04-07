package commands

import (
	"bot/pkg/sql"
	"bot/pkg/util"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// driver stats
func stats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, "Not a valid target", true)
		return
	}

	id, err := strconv.Atoi(user.ID)
	util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)

	if sql.CheckDriver(id) {
		driver := sql.GetDriver(id)
		msg := driver.Out()

		respond(s, i, msg, false)
	} else {
		respond(s, i, "Driver does not exist!", true)
	}
}

// wins
func wins(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, "Not a valid target", true)
		return
	}

	id, err := strconv.Atoi(user.ID)
	util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)

	if sql.CheckDriver(id) {
		wins := sql.GetWins(id)
		msg := fmt.Sprintf(
			"``%s has %d wins.``",
			user.Username,
			wins,
		)

		respond(s, i, msg, true)
	} else {
		respond(s, i, "Driver does not exist!", true)
	}
}

// podiums
func podiums(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, "Not a valid target", true)
		return
	}

	id, err := strconv.Atoi(user.ID)
	util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)

	if sql.CheckDriver(id) {
		wins := sql.GetPodiums(id)
		msg := fmt.Sprintf(
			"``%s has %d podiums.``",
			user.Username,
			wins,
		)

		respond(s, i, msg, true)
	} else {
		respond(s, i, "Driver does not exist!", true)
	}
}

// points
func points(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, "Not a valid target", true)
		return
	}

	id, err := strconv.Atoi(user.ID)
	util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)

	if sql.CheckDriver(id) {
		wins := sql.GetPoints(id)
		msg := fmt.Sprintf(
			"``%s has %.2f points.``",
			user.Username,
			wins,
		)

		respond(s, i, msg, true)
	} else {
		respond(s, i, "Driver does not exist!", true)
	}
}
