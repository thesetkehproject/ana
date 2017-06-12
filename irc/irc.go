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

func Connect(config configuration.Container) {
	connection := irclib.IRC(config.IrcConfig.BotNick, config.AnaCommon.BotUser)
	connection.UseTLS = true
	connection.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	connection.AddCallback("001", func(e *irclib.Event) { connection.Join(config.IrcConfig.Channel) })
	connection.AddCallback("366", func(e *irclib.Event) { })
	connection.AddCallback("PRIVMSG", func(event *irclib.Event) {
		fmt.Printf("%v: %v\n", event.Nick, event.Message())
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.AnaCommon.LogDir), fmt.Sprintf("%v: %v", event.Nick, event.Message()))
	})
	err := connection.Connect(config.IrcConfig.Server)

	if err != nil {
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.AnaCommon.LogDir), fmt.Sprintf("%v\n", err))
		connection.Quit()
	}
	connection.Loop()
}

//func SendNotification(target, message string, logdir string) {
//	if !connection.Connected() {
//		if err := Connect(); err != nil {
//			logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", logdir), fmt.Sprintf("%v\n", err))
//		}
//	}
//	connection.Notice(target, message)
//	logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", logdir),
//		fmt.Sprintf("Target: %s, message: %s\n", target, message))
//}
