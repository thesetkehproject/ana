package configuration

import (
	"encoding/json"
	"log"
)

type Container struct {
	Common struct {
		BotUser    string       `json:"BotUser"`
		Admins     []string     `json:"Admins"`
		Trigger    string       `json:"Trigger"`
		LogDir     string       `json:"LogDir"`
		WikiLink   string       `json:"WikiLink"`
		Homepage   string       `json:"Homepage"`
		Forums     string       `json:"Forums"`
		WeatherKey string       `json:"WeatherKey"`
	}  `json:"anaCommon"`
	ircConfig struct {
		Server     string       `json:"Server"`
		Channel    string       `json:"Channel"`
		BotNick    string       `json:"BotNick"`
	}   `json:"ircConfig"`
}

func doConfig() Container {
	jStr := `
    {
        "anaCommon": {
            "BotUser": "Eliza-BOT-Ana",
            "Admins": ["darthlukan", "setkeh"],
            "Trigger": "!",
            "LogDir": "/tmp/ana/",
            "WikiLink": "https://github.com/thesetkehproject",
            "Homepage": "https://github.com/thesetkehproject",
            "Forums": "https://github.com/thesetkehproject",
            "WeatherKey": "AMSL"
        },
        "ircConfig": {
            "Server": "irc.freenode.net:6667",
            "Channel": "#thesetkehproject",
            "BotNick": "Eliza-BOT-Ana"
        }
    }
    `

	var cont Container
	if err := json.Unmarshal([]byte(jStr), &cont); err != nil {
		log.Fatal(err)
	}

	return cont
}