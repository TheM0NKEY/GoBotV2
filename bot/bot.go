package bot

import (
	"GoBotV2/config"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Run() {
	//Create Bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	BotID = user.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		return
	}

	defer goBot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to Exit")
	<-stop
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "Hi" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi Back")
	}
}
