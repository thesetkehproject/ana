// Ana
package main

import (
	"encoding/json"
	"fmt"
	"github.com/thesetkehproject/ana/logger"
	"github.com/thoj/go-ircevent"
	"os"
	"strings"
)

const delay = 40

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

// ParseCmds takes PRIVMSG strings containing a preceding bang "!"
// and attempts to turn them into an ACTION that makes sense.
// Returns a msg string.
func ParseCmds(cmdMsg string, config *Config) string {
	var (
		msg      string
		msgArray []string
		cmdArray []string
	)

	cmdArray = strings.SplitAfterN(cmdMsg, config.Trigger, 2)

	if len(cmdArray) > 0 {
		msgArray = strings.SplitN(cmdArray[1], " ", 2)
	}

	switch {
	case strings.Contains(msgArray[0], "help"):
		msg = HelpCmd(config.Trigger)
	case strings.Contains(msgArray[0], "wiki"):
		msg = WikiCmd(config)
	case strings.Contains(msgArray[0], "homepage"):
		msg = HomePageCmd(config)
	case strings.Contains(msgArray[0], "forums"):
		msg = ForumCmd(config)
	default:
		msg = ""
	}
	return msg
}

// AddCallbacks is a single function that does what it says.
// It's merely a way of decluttering the main function.
func AddCallbacks(conn *irc.Connection, config *Config) {
	log := fmt.Sprintf("%s%s", config.LogDir, config.Channel)

	conn.AddCallback("001", func(e *irc.Event) {
		conn.Join(config.Channel)
	})

	conn.AddCallback("JOIN", func(e *irc.Event) {
		if e.Nick == config.BotNick {
			conn.Privmsg(config.Channel, "Hello everybody, I'm a bot")
			logger.LogDirCheck(config.LogDir)
			logger.LogFile(config.LogDir + e.Arguments[0])
		}
		message := fmt.Sprintf("%s has joined", e.Nick)
		go logger.IRCChannelLogger(log, e.Nick, message)
	})
	conn.AddCallback("PART", func(e *irc.Event) {
		message := fmt.Sprintf("has parted (%s)", e.Message())
		nick := fmt.Sprintf("%s@%s", e.Nick, e.Host)
		go logger.IRCChannelLogger(log, nick, message)
	})
	conn.AddCallback("QUIT", func(e *irc.Event) {
		message := fmt.Sprintf("has quit (%v)", e.Message)
		nick := fmt.Sprintf("%s@%s", e.Nick, e.Host)
		go logger.IRCChannelLogger(log, nick, message)
	})

	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		var response string
		message := e.Message()
		if strings.Contains(message, config.Trigger) && strings.Index(message, config.Trigger) == 0 {
			response = ParseCmds(message, config)
		}

		if strings.Contains(message, fmt.Sprintf("%squit", config.Trigger)) {
			QuitCmd(config.Admins, e.Nick)
		}

		if len(response) > 0 {
			conn.Privmsg(config.Channel, response)
		}

		if len(message) > 0 {
			if e.Arguments[0] != config.BotNick {
				go logger.IRCChannelLogger(log, e.Nick+": ", message)
			}
		}
	})
}

func main() {

	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println("Couldn't read config file, dying...")
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	decoder.Decode(&config)

	conn := irc.IRC(config.BotNick, config.BotUser)
	err = conn.Connect(config.Server)

	if err != nil {
		fmt.Println("Failed to connect.")
		panic(err)
	}

	AddCallbacks(conn, config)
	conn.Loop()
}
