package main

import (
	"log"
	"os"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/wafellofazztrack/lectronic-backend/command"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	err := command.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
