package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/santosh/discord-ping/config"
)

var BotId string
var goBot *discordgo.Session

func Start() {

	goBot, err1 := discordgo.New("Bot " + config.Token)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	u, err2 := goBot.User("@me")

	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)
	err3 := goBot.Open()
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}
	fmt.Println("Bot is running!")

}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// if m.Author.ID == BotId {
	// 	return
	// }
	// if m.content contains botid (Mentions) and "ping" then send "pong!"
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}
	if m.Content == "what is my name" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello Santosh!")
	}
}
