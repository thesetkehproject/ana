// Ana
package main

import (
	Config "github.com/thesetkehproject/ana/configuration"
	"fmt"
)

func main() {
	ConfigContainer := Config.DoConfig("config.json")

	fmt.Println("Bot User: %v", ConfigContainer.AnaCommon.BotUser)

	if ConfigContainer.IrcConfig.Server != "" {
		fmt.Println(ConfigContainer.IrcConfig)
	} else {
		fmt.Println("No IRC Configuration!")
	}
}
