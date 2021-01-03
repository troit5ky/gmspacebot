package main

import (
	"log"

	"./lib/Bot"
	"./lib/Config"
)

func main() {
	log.Println("Init bot config...")
	log.Println(Config.Init())

	log.Println("Init Bot...")
	Bot.Init()
}
