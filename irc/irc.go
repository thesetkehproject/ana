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
	connection := irclib.IRC(config.IrcConfig.BotNick, config.AnaCommon.BotUser)
	connection.UseTLS = true
	connection.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	connection.AddCallback("001", func(e *irclib.Event) {
		connection.Notice(config.IrcConfig.Channel, message)
		connection.Quit()
	})

	connection.AddCallback("366", func(e *irclib.Event) { })

	err := connection.Connect(config.IrcConfig.Server)
	connection.Loop()

	if err != nil {
		fmt.Sprintf("%v\n", err)
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.AnaCommon.LogDir), fmt.Sprintf("%v\n", err))
		connection.Quit()
	}
}
