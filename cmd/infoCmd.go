package cmd

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func UserInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate, id string) {
	user, err := s.GuildMember(m.GuildID, id)
	if err != nil{
		embed := &discordgo.MessageEmbed{
			Description: "Invalid User.",
			Title: "Error",
			Color: 15158332,
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return
	}

	fieldName := &discordgo.MessageEmbedField{
		Name:   "Name",
		Value:  user.User.Username,
		Inline: true,
	}

	fieldTag := &discordgo.MessageEmbedField{
		Name:   "Tag",
		Value:  user.User.Discriminator,
		Inline: true,
	}

	var nickname string
	if user.Nick == ""{
		nickname = "None"
	}

	fieldNick := &discordgo.MessageEmbedField{
		Name:   "Nickname",
		Value:  nickname,
		Inline: true,
	}

	fieldID := &discordgo.MessageEmbedField{
		Name:   "ID",
		Value:  user.User.ID,
		Inline: true,
	}

	fieldBot := &discordgo.MessageEmbedField{
		Name:   "Bot",
		Value:  strconv.FormatBool(user.User.Bot),
		Inline: true,
	}

	value, _ := discordgo.SnowflakeTimestamp(id)
	timeFromCreation := value.String()

	fieldCreated := &discordgo.MessageEmbedField{
		Name:   "Account Created",
		Value:  timeFromCreation[:19],
		Inline: true,
	}

	joinDate := string(user.JoinedAt)
	joinFilter := strings.Replace(joinDate, "T", " ", -1)

	fieldJoined := &discordgo.MessageEmbedField{
		Name:   "Joined in this server",
		Value:  joinFilter[:19],
		Inline: true,
	}

	var roleList []string
	guildRoles, _ := s.GuildRoles(m.GuildID)
	for _, x:= range user.Roles{
		for _, r := range guildRoles{
			if r.ID == x && r.Color != 3092790{
				role := fmt.Sprintf("<@&%v>", x)
				roleList = append(roleList, role)
			}
		}
	}
	roles := strings.Join(roleList, " ")

	fieldRoles := &discordgo.MessageEmbedField{
		Name:   "Roles",
		Value:  roles,
		Inline: false,
	}


	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf("Info about: %v\n", user.Mention()),
		Title: "UserInfo",
		Fields: []*discordgo.MessageEmbedField{
			fieldName,
			fieldTag,
			fieldNick,
			fieldID,
			fieldBot,
			fieldCreated,
			fieldJoined,
			fieldRoles,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: user.User.AvatarURL("")},
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func BotInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate, upTime string) {

	embedAuthor := &discordgo.MessageEmbedAuthor{
		Name: fmt.Sprintf("%v#%v",s.State.User.Username,s.State.User.Discriminator),
		IconURL: s.State.User.AvatarURL(""),
	}

	developer, _ := s.User("173569203977060353")
	embedDeveloper := &discordgo.MessageEmbedField{
		Name: "Developer:",
		Value: fmt.Sprintf("%v#%v",developer.Username,developer.Discriminator),
		Inline: true,
	}


	embedUptime := &discordgo.MessageEmbedField{
		Name: "Uptime:",
		Value: upTime,
		Inline: true,
	}

	var users int
	var guilds int
	for _, guild := range s.State.Guilds {
		guilds += 1
		users += guild.MemberCount
	}


	embedUsers := &discordgo.MessageEmbedField{
		Name: "Users:",
		Value: strconv.Itoa(users),
		Inline: true,
	}

	embedGuilds := &discordgo.MessageEmbedField{
		Name: "Guilds:",
		Value: strconv.Itoa(guilds),
		Inline: true,
	}

	embedLinks := &discordgo.MessageEmbedField{
		Name: "Links:",
		Value: "[Donations](https://ko-fi.com/polliog) | [Github](https://github.com/Polliog)",
		Inline: false,
	}

	embed := &discordgo.MessageEmbed{
		Author: embedAuthor,
		Fields: []*discordgo.MessageEmbedField{
			embedDeveloper,
			embedUptime,
			embedUsers,
			embedGuilds,
			embedLinks,
		},
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}