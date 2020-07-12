package cmd

import (
	"github.com/bwmarrin/discordgo"
)

func UnbanCommand(s *discordgo.Session, m *discordgo.MessageCreate, id string) {

	permission, err := s.State.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil{
		embed := &discordgo.MessageEmbed{
			Description: "Error getting user permissions.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	if permission&discordgo.PermissionBanMembers == 0{
		embed := &discordgo.MessageEmbed{
			Description: "you don't have the permission for unban members.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	bans, err := s.GuildBans(m.GuildID)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Description: "Error getting guild bans.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	if id == m.Author.ID{
		embed := &discordgo.MessageEmbed{
			Description: "You can't unban yourself...",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	exists := false
	for _, ban := range bans {
		if ban.User.ID == id {
			exists = true
		}
	}

	if !exists {
		embed := &discordgo.MessageEmbed{
			Description: "The selected member isn't banned",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	err = s.GuildBanDelete(m.GuildID, id)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Description: "Failed to unban the user.",
			Title:       "Error",
			Color:       15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, "The user can rejoin the server.")
	return
}
