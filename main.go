package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/dragonator/coach-ai-assignment/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
