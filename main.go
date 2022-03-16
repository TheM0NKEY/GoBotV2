package main

import (
	"GoBotV2/bot"
	"GoBotV2/config"
	"log"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Run()
	//<-make(chan struct{})
	return
}
