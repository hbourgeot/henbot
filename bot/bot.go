package bot

import (
	"log"
	"os"
	"time"

	discordgo2 "github.com/bwmarrin/discordgo"
	"github.com/hbourgeot/henbot/config"
)

var (
	BotID string
	goBot *discordgo2.Session
)

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
	goBot.AddHandler(welcomeHandler)
	if err = goBot.Open(); err != nil {
		log.Fatal(err)
		return
	}
}

func messageHandler(s *discordgo2.Session, m *discordgo2.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	content := m.Author.Mention() + "\n" // for show any content on the message

	switch m.Content { // we switch on the message sent by anyone
	case "/help":
		file, err := os.ReadFile("./help.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		content += string(file)
		break
	case "/github":
		content += "Click on the URL for see my github https://github.com/hbourgeot"
		break
	case "/portfolio":
		content += "Click on the URL for see my portfolio https://www.henrry.online"
		break
	default:
		return
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, content) // the bot sen the message
}

func welcomeHandler(s *discordgo2.Session, m *discordgo2.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	hour, minute, second := m.Member.JoinedAt.Local().Clock()

	if hour == time.Now().Local().Hour() && minute == time.Now().Local().Minute() && second <= time.Now().Local().Second() {
		content := "Welcome " + m.Author.Mention() + "!\n Please read the rules and use the /help command for know what I can do"
		_, _ = s.ChannelMessageSend(m.ChannelID, content)
	}
}
