package bot

import (
	discordgo2 "github.com/bwmarrin/discordgo"
	"github.com/hbourgeot/henbot/config"
	"log"
	"os"
)

var BotID string
var goBot *discordgo2.Session

func Run() {
	goBot, err := discordgo2.New("Bot " + config.Token) // we make new bot
	if err != nil {
		log.Fatal(err)
		return
	}

	user, err := goBot.User("@me") // create new user
	if err != nil {
		log.Fatal(err)
		return
	}

	BotID = user.ID // we assign the user.ID to the BotID
	goBot.AddHandler(messageHandler)
	if err = goBot.Open(); err != nil {
		log.Fatal(err)
		return
	}
}

func messageHandler(s *discordgo2.Session, m *discordgo2.MessageCreate) {
	var content string //for show any content on the message
	if m.Author.ID == BotID {
		return
	}
	switch m.Content { //we switch on the message sent by anyone
	case "/help":
		file, err := os.ReadFile("./help.txt")
		if err != nil {
			log.Fatal(err)
			return
		}

		content = string(file)
		break
	case "/github":
		content = "Click on the URL for see my github https://github.com/hbourgeot"
		break
	case "/portfolio":
		content = "Click on the URL for see my portfolio https://portfolio-hb.herokuapp.com"
		break
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, content) //the bot sen the message

}