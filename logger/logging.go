package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func IRCChannelLogger(logfile string, nick string, message string) {
	STime := time.Now().UTC().Format(time.ANSIC)
	prefile := strings.Replace(logfile, "#", "", 1)
	file := fmt.Sprintf("%s.log", prefile)

	//Open the file for writing With Append Flag to create file persistence
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	n, err := io.WriteString(f, fmt.Sprintf("%v > %v: %v\n", STime, nick, message))
	if err != nil {
		fmt.Println(n, err)
	}
}

func LogDirCheck(logdir string) {
	if _, err := os.Stat(logdir); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s, creating.\n", logdir)
		os.Mkdir(logdir, 0777)
	} else {
		fmt.Printf("%s exists.\n", logdir)
	}
}

func LogFile(logfile string) {
	prelog := strings.Replace(logfile, "#", "", 1)
	file := fmt.Sprintf("%s.log", prelog)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("Log file '%s' does not Exist. Creating.\n", file)
		os.Create(file)
		fmt.Printf("Log file '%s' created.\n", file)
	} else {
		fmt.Printf("Log file '%s' exists, reusing.\n", file)
	}
}
