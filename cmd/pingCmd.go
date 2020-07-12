package cmd

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func PingCommand(s* discordgo.Session, m* discordgo.MessageCreate) {
	messageTime, err := m.Timestamp.Parse()
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Description: "Error parsing timestamp",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	message := fmt.Sprintf("**Ping**: %v\n", time.Since(messageTime).Round(time.Millisecond).String())
	s.ChannelMessageSend(m.ChannelID, message)
	return

}
