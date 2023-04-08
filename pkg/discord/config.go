package discord

import (
	"bot/pkg/util"
	"encoding/json"
	"os"
	"sync"
)

type ConfigData struct {
	Token  string
	Status string
}

var lock = &sync.Mutex{}

var instance *ConfigData

func Config() *ConfigData {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		instance = getConfig()
	}

	return instance
}

func getConfig() *ConfigData {
	var cd ConfigData

	file, err := os.ReadFile("configs/discord.json")
	util.CheckErrMsg(err, "Config read failure")

	err = json.Unmarshal(file, &cd)
	util.CheckErrMsg(err, "Config unmarshall error")

	return &cd
}
