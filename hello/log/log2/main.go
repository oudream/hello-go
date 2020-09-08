package main

import (
	"log"
	"os"
)

var logFile *os.File

func initSetLogOutput() {
	var err error
	logFile, err = os.OpenFile("errors.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(logFile)
	log.Println("This is a test log entry")
}

func closeLogOutput() {
	_ = logFile.Close()
}

func main() {
	initSetLogOutput()
	defer closeLogOutput()
	log.Println("begin:")
	log.Println("end.")
}
