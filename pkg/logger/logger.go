package logger

import (
	"fmt"
	"log"
	"os"
)

// LogType holds the log level.
type LogType int

const (
	// INFO logs Info, Warnings and Errors
	INFO LogType = iota // 0 :
	// WARNING logs Warning and Errors
	WARNING LogType = iota //1
	// ERROR just logs Errors
	ERROR LogType = iota // 2
)

//InitializeLogger creates Logger by file
func InitializeLogger(logFile string) error {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.SetOutput(file)
	return nil

}

//Log logs what you want
func Log(message string, logType LogType) {

	if logType == INFO {
		log.Println("INFO:", message)
		fmt.Println("INFO:", message)
	}
	if logType == WARNING {
		log.Println("WARNING:", message)
		fmt.Println("WARNING:", message)
	}
	if logType == ERROR {
		log.Println("ERROR:", message)
		fmt.Println("ERROR:", message)
	}

}
