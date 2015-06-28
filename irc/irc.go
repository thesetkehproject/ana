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

// AddCallbacks is a single function that does what it says.
// It's merely a way of decluttering the main function.
func AddCallbacks(conn *irclib.Connection, config *Config) {
	log := fmt.Sprintf("%s%s", config.LogDir, config.Channel)

	conn.AddCallback("001", func(e *irclib.Event) {
		conn.Join(config.Channel)
	})

	conn.AddCallback("JOIN", func(e *irclib.Event) {
		if e.Nick == config.BotNick {
			conn.Privmsg(config.Channel, "Hello everybody, I'm a bot")
			logger.LogDirCheck(config.LogDir)
			logger.LogFile(config.LogDir + e.Arguments[0])
		}
		message := fmt.Sprintf("%s has joined", e.Nick)
		go logger.IRCChannelLogger(log, e.Nick, message)
	})
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
	AddCallbacks(connection, config)
	connection.Loop()
}
