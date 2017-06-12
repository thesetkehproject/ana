// Ana
package main

import (
	"github.com/thesetkehproject/ana/configuration"
	"fmt"
)

func main() {
	config := configuration.Container{}.Common

	fmt.Printf(config.Botuser)
}
