package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	InfoLogger    *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	WarningLogger = log.New(file, "Warning: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "Error: ", log.LstdFlags|log.Lshortfile)
	InfoLogger = log.New(file, "Info: ", log.LstdFlags|log.Lshortfile)
}

func main() {
	InfoLogger.Println("This is some info")
	WarningLogger.Println("This is a warning")
	ErrorLogger.Println("this is an error")
}
