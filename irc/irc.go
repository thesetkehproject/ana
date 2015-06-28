package irc

import (
	"encoding/json"
	"fmt"
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

func SendNotification(target, message string) {
	connection.Notice(target, message)
	// TODO: Logger needs to be able to log more than just a channel, but notices as well
	fmt.Printf("Wrote message: '%s' to target '%s'", message, target)
}

func init() {
	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println("Couldn't read config file, dying...")
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config = &Config{}
	decoder.Decode(&config)

	connection := irclib.IRC(config.BotNick, config.BotUser)
	err = connection.Connect(config.Server)

	if err != nil {
		fmt.Println("Failed to connect.")
		panic(err)
	}
}

func Client() {
	connection.Loop()
}
