package main

import (
	"github.com/hbourgeot/henbot/bot"
	"github.com/hbourgeot/henbot/config"
	"log"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Run()
	<-make(chan struct{}) //Channel for goroutines
	return
}