package cmd

import (
	"../util"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func KickCommand(s *discordgo.Session, m *discordgo.MessageCreate, id string, reason string) {
	author, _ :=  s.GuildMember(m.GuildID, m.Author.ID)

	guild, _ := s.Guild(m.GuildID)

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

	if permission&discordgo.PermissionKickMembers == 0{
		embed := &discordgo.MessageEmbed{
			Description: "you don't have the permission for kick members.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	if id == m.Author.ID{
		embed := &discordgo.MessageEmbed{
			Description: "You can't kick yourself.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	victim, err := s.GuildMember(m.GuildID, id)
	if err != nil || victim == nil {
		embed := &discordgo.MessageEmbed{
			Description: "Invalid User.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	if util.RoleGerarchyDifference(victim, author, guild) >= 0{
		embed := &discordgo.MessageEmbed{
			Description: "You can only kick members with lower permissions than yours.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	embedDM := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("You are kicked from %v", guild.Name),
		Description: fmt.Sprintf("`Reason`: %v\n\n`Author`: %v", reason, m.Author.Username),
	}
	channelDM, _ := s.UserChannelCreate(victim.User.ID)
	_, _ = s.ChannelMessageSendEmbed(channelDM.ID, embedDM)


	err = s.GuildMemberDeleteWithReason(m.GuildID, id, reason)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Description: fmt.Sprintf("Failed to kick %v.\n", victim.Mention()),
			Title:       "Error",
			Color:       15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}
	message := fmt.Sprintln(victim.Mention(), "is now kicked from this server")
	_, _ = s.ChannelMessageSend(m.ChannelID, message)
	return
}
