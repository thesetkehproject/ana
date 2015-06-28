package main

import (
	"fmt"
	"os"
)

func HelpCmd(trigger string) string {
	return fmt.Sprintf("Commands: %shelp, %sddg/search, %sconvtemp, %scakeday, %sVERB. Admins only: %squit\n",
		trigger, trigger, trigger, trigger, trigger, trigger)
}

func WikiCmd(config *Config) string {
	return fmt.Sprintf("(Channel Wiki)[ %s ]\n", config.WikiLink)
}

func HomePageCmd(config *Config) string {
	return fmt.Sprintf("(Channel Homepage)[ %s ]\n", config.Homepage)
}

func ForumCmd(config *Config) string {
	return fmt.Sprintf("(Channel Forums)[ %s ]\n", config.Forums)
}

func QuitCmd(admins []string, user string) {
	for _, admin := range admins {
		if user == admin {
			os.Exit(0)
		}
	}
}
