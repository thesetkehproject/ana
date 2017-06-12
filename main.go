// Ana
package main

import (
	Config "github.com/thesetkehproject/ana/configuration"
	"fmt"
	"github.com/thesetkehproject/ana/irc"
	"github.com/thesetkehproject/ana/logger"
)

func main() {
	ConfigContainer := Config.DoConfig("config.json")

	if ConfigContainer.IrcConfig.Server != "" {
		logger.GenericLogger(ConfigContainer.AnaCommon.LogDir, "IRC Configuration Found! Running Connection Command")
		irc.SendIrcNotice(ConfigContainer, "New Message Test Function From Main Function.")
	} else {
		fmt.Println("No IRC Configuration!")
	}
}
