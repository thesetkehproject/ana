package irc

import (
	"encoding/json"
	"fmt"
	"github.com/thesetkehproject/ana/logger"
	irclib "github.com/thoj/go-ircevent"
	"os"
)

var (
	config     *Config
	connection *irclib.Connection
)

type Config struct {
	Admins     []string
	Server     string
	Channel    string
	BotUser    string
	BotNick    string
	Trigger    string
	WeatherKey string
	LogDir     string
	WikiLink   string
	Homepage   string
	Forums     string
}

func connect() error {
	connection := irclib.IRC(config.BotNick, config.BotUser)
	err := connection.Connect(config.Server)

	if err != nil {
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.LogDir), fmt.Sprintf("%v\n", err))
		return err
	}
	return nil
}

func init() {
	file, err := os.Open("config.json")

	if err != nil {
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.LogDir), fmt.Sprintf("%v\n", err))
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config = &Config{}
	decoder.Decode(&config)

	err = connect()
	if err != nil {
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.LogDir), fmt.Sprintf("%v\n", err))
		panic(err)
	}
}

func SendNotification(target, message string) {
	if !connection.Connected() {
		if err := connect(); err != nil {
			logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.LogDir), fmt.Sprintf("%v\n", err))
		}
	}
	connection.Notice(target, message)
	logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.LogDir),
		fmt.Sprintf("Target: %s, message: %s\n", target, message))
}
