package main

import (
	"log"

	"github.com/hbourgeot/henbot/bot"
	"github.com/hbourgeot/henbot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Run()
	<-make(chan struct{}) // Channel for goroutines
	return
}
