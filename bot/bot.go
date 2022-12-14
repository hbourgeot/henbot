package bot

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	discordgo2 "github.com/bwmarrin/discordgo"
	"github.com/hbourgeot/henbot/calc"
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
	goBot.AddHandler(calcHandler)

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
	case "/help", "/ayuda", "/HELP", "/AYUDA":
		file, err := os.ReadFile("./help.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		content += string(file)
		break
	case "/sociales", "/SOCIALES":
		file, err := os.ReadFile("./socials.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		content += string(file)
		break
	case "/github", "/GITHUB", "/gh":
		content += "Haz click en el link para ver mi github: https://github.com/hbourgeot"
		break
	case "/portfolio", "/portafolio", "/PORTAFOLIO", "/PORTFOLIO":
		content += "Haz click en el link para ver mi portafolio: https://www.henrry.online"
		break
	case "/bot", "/BOT":
		content += "Aquí puedes ver el repo del bot: https://github.com/hbourgeot/henbot"
	default:
		return
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, content) // the bot sen the message
}

func calcHandler(s *discordgo2.Session, m *discordgo2.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	var (
		res     int
		err     error
		content string
	)

	operation := strings.Split(m.Content, " ")
	if strings.Contains(m.Content, "/calc-sum") {
		res, err = calc.Add(operation)
		if err != nil {
			return
		}

		content = m.Author.Mention() + " la suma es igual a: "
	} else if strings.Contains(m.Content, "/calc-sub") {
		res, err = calc.Sub(operation)
		if err != nil {
			return
		}

		content = m.Author.Mention() + " la resta es igual a: "
	} else if strings.Contains(m.Content, "/calc-mul") {
		res, err = calc.Mult(operation)
		if err != nil {
			return
		}

		content = m.Author.Mention() + " la multiplicación es igual a: "
	} else if strings.Contains(m.Content, "/calc-div") {
		res, err = calc.Div(operation)
		if err != nil {
			return
		}

		content = m.Author.Mention() + " la división es igual a: "
	} else if strings.Contains(m.Content, "/calc-mod") {
		res, err = calc.Mod(operation)
		if err != nil {
			return
		}

		content = m.Author.Mention() + " el módulo es igual a: "
	} else if strings.Contains(m.Content, "/calc-perc") {
		percentage, err := calc.Perc(operation)
		if err != nil {
			return
		}

		content = m.Author.Mention() + " el " + fmt.Sprint(operation[1]) + "% de " + fmt.Sprint(operation[2]) + " es igual a: "
		_, _ = s.ChannelMessageSend(m.ChannelID, content+fmt.Sprint(percentage))
		return
	} else {
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, content+fmt.Sprint(res))
}

func welcomeHandler(s *discordgo2.Session, m *discordgo2.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	hour, minute, second := m.Member.JoinedAt.Local().Clock()

	if hour == time.Now().Local().Hour() && minute == time.Now().Local().Minute() && second <= time.Now().Local().Second() {
		content := "Bienvenido a mi servidor " + m.Author.Mention() + "!\n Por favor lee las reglas del servidor y si quieres usa el comando /help para conocer lo que puedo hacer"
		_, _ = s.ChannelMessageSend("1017897619843977278", content)
	}
}
