package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/thesetkehproject/ana/logger"
)

var (
	cont Container
)

type Container struct {
	AnaCommon struct {
		BotUser    string       `json:"BotUser"`
		Admins     []string     `json:"Admins"`
		Trigger    string       `json:"Trigger"`
		LogDir     string       `json:"LogDir"`
		WikiLink   string       `json:"WikiLink"`
		Homepage   string       `json:"Homepage"`
		Forums     string       `json:"Forums"`
		WeatherKey string       `json:"WeatherKey"`
	}  `json:"AnaCommon"`
	IrcConfig struct {
		Server     string       `json:"Server"`
		Channel    string       `json:"Channel"`
		BotNick    string       `json:"BotNick"`
	}   `json:"IrcConfig"`
}

func DoConfig(filePath string) Container {
	file, err := os.Open(filePath)

	if err != nil {
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", cont.AnaCommon.LogDir), fmt.Sprintf("%v\n", err))
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&cont)

	return cont
}