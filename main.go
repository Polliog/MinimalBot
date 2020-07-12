package main

import (
	"./cmd"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	Token string
	startTime time.Time
)



func init()  {
	flag.StringVar(&Token, "t", "","Bot Token")
	flag.Parse()
	startTime = time.Now()
}


func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error when creating the discord session:\n", err)
		return
	}



	dg.AddHandler(commandHandler)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error when opening the discord connection\n", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()

}


func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content
	commandContent := strings.Split(content, " ")
	command := commandContent[0]

	switch {
	case command == "m!say":
		var message3Filter string
		if len(commandContent) < 2{
			embed := &discordgo.MessageEmbed{
				Title: "Error",
				Description: "Invalid Command Usage:\n`m!say [message]`",
				Color: 15158332,
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			return
		} else {
			message1Filter := strings.Replace(content, "m!say", "", -1)
			message2Filter := strings.Replace(message1Filter, "@everyone", "`everyone`", -1)
			message3Filter = strings.Replace(message2Filter, "@here", "`here`", -1)
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, message3Filter)
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
	case command == "m!help":
		cmd.HelpCommand(s, m)
	case command == "m!ban" && len(commandContent[1]) > 0:
		reasonFilter := strings.Replace(content, "m!ban", "", -1)
		reason := strings.Replace(reasonFilter, commandContent[1], " ", -1)
		cmd.BanCommand(s, m, commandContent[1], reason)
	case command == "m!ping":
		cmd.PingCommand(s, m)
	case command == "m!unban" && len(commandContent[1]) > 0:
		cmd.UnbanCommand(s, m, commandContent[1])
	case command == "m!kick" && len(commandContent[1]) > 0:
		reasonFilter := strings.Replace(content, "m!ban", "", -1)
		reason := strings.Replace(reasonFilter, commandContent[1], " ", -1)
		cmd.BanCommand(s, m, commandContent[1], reason)
	case command == "m!userinfo":
		var userID string
		if len(commandContent) < 2{
			userID = m.Author.ID
		} else {
			userID = commandContent[1]
		}
		cmd.UserInfoCommand(s, m, userID)
	case command == "m!info":
		timeDifference := humanize.Time(startTime)
		cmd.BotInfoCommand(s, m, timeDifference)
	}
}