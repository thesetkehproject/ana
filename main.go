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
	logfile := fmt.Sprintf("%v/%v", ConfigContainer.AnaCommon.LogDir, ConfigContainer.AnaCommon.LogFile)

	fmt.Printf("Checking to make sure Logfile %v Exists in Log Directory %v\n", ConfigContainer.AnaCommon.LogFile, ConfigContainer.AnaCommon.LogDir)
	logger.LogDirCheck(ConfigContainer.AnaCommon.LogDir)
	logger.LogFileCheck(ConfigContainer.AnaCommon.LogDir, ConfigContainer.AnaCommon.LogFile)
	fmt.Println("Logger Is Initialized")
	logger.GenericLogger(logfile, "Logger Is Initialized")

	if ConfigContainer.IrcConfig.Server != "" {
		logger.GenericLogger(logfile, "IRC Configuration Found! Running Connection Command")
		irc.SendIrcNotice(ConfigContainer, "New Message Test Function From Main Function.")
	} else {
		fmt.Println("No IRC Configuration!")
		logger.GenericLogger(logfile, "No IRC Configuration!")
	}
}
