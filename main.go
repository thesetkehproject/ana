// Ana
package main

import (
	"github.com/ana/configuration"
	"fmt"
)

func main() {
	config := configuration.Container{}.Common

	fmt.Printf(config.Botuser)
}
