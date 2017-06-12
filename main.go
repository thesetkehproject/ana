// Ana
package main

import (
	Config "github.com/thesetkehproject/ana/configuration"
	"fmt"
)

func main() {
	Container := Config.DoConfig("config.json")

	fmt.Printf("Bot User: %v", Container.AnaCommon.BotUser)
}
