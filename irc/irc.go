package irc

import (
	"fmt"
	"github.com/thesetkehproject/ana/logger"
	irclib "github.com/thoj/go-ircevent"
	"github.com/thesetkehproject/ana/configuration"
	"crypto/tls"
)

var (
	config configuration.Container
	connection *irclib.Connection
)

func Connect() error {
	connection := irclib.IRC(config.IrcConfig.BotNick, config.AnaCommon.BotUser)
	connection.UseTLS = true
	connection.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	fmt.Println("ircServer: %v", config.IrcConfig.Server)
	err := connection.Connect(config.IrcConfig.Server)

	if err != nil {
		logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.AnaCommon.LogDir), fmt.Sprintf("%v\n", err))
		return err
	}
	return nil
}

func SendNotification(target, message string) {
	if !connection.Connected() {
		if err := Connect(); err != nil {
			logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.AnaCommon.LogDir), fmt.Sprintf("%v\n", err))
		}
	}
	connection.Notice(target, message)
	logger.GenericLogger(fmt.Sprintf("%s/ana-irc.log", config.AnaCommon.LogDir),
		fmt.Sprintf("Target: %s, message: %s\n", target, message))
}
