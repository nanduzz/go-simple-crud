package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ErrorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
var InfoLog = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
var WarnrLog = log.New(os.Stdout, "WARN : ", log.Ldate|log.Ltime|log.Lshortfile)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
