package commands

import (
	"bot/internal/racing"
	"bot/pkg/util"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// qualifying
func qualifying(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, invalidTarget, true)
		return
	}

	series := i.ApplicationCommandData().Options[1].StringValue()
	race := i.ApplicationCommandData().Options[2].StringValue()
	season := i.ApplicationCommandData().Options[3].StringValue()
	year := i.ApplicationCommandData().Options[4].StringValue()

	userId, err := strconv.Atoi(user.ID)
	util.CheckErrMsg(err, "Conversion error occurred converted User Id: "+user.ID)
	qualiResultId := racing.GenQualiId(series, race, season, year)

	c := getDatabase()
	if c.Conn.CheckQualifyingResult(qualiResultId, userId) {
		qualiResult := c.Conn.QualifyingResult(qualiResultId, userId)
		msg := fmt.Sprintf(
			">>> **Qualifying result for %s**: \n%s",
			user.Username,
			qualiResult.Out(),
		)
		respond(s, i, msg, false)
	} else {
		respond(s, i, invalidQuali, true)
	}
}
