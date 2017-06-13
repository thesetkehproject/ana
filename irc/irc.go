package irc

import (
	"fmt"
	"github.com/thesetkehproject/ana/logger"
	irclib "github.com/thoj/go-ircevent"
	"crypto/tls"
	"github.com/thesetkehproject/ana/configuration"
)

var (
	connection *irclib.Connection
)

func SendIrcNotice(config configuration.Container, message string) {
	logfile := fmt.Sprintf("%v/%v", config.AnaCommon.LogDir, config.AnaCommon.LogFile)
	connection := irclib.IRC(config.IrcConfig.BotNick, config.AnaCommon.BotUser)
	connection.UseTLS = true
	connection.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	connection.AddCallback("001", func(e *irclib.Event) {
		logger.GenericLogger(logfile, fmt.Sprintf("Sending Notice: %v To the Channel %v\n", message, config.IrcConfig.Channel))
		connection.Notice(config.IrcConfig.Channel, message)
		connection.Quit()
	})

	connection.AddCallback("366", func(e *irclib.Event) { })

	err := connection.Connect(config.IrcConfig.Server)
	connection.Loop()

	if err != nil {
		fmt.Printf("%v\n", err)
		logger.GenericLogger(logfile, fmt.Sprintf("%v\n", err))
		connection.Quit()
	}
}
