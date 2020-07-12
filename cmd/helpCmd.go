package cmd

import "github.com/bwmarrin/discordgo"

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	embed := &discordgo.MessageEmbed{
		Title: "Help Command",
		Description:
			"`Help`: Shows this menu;\n" +
			"`Say [message]`: The bot will repeat your message;\n" +
			"`Ping`: Shows the ping of the bot;\n" +
			"`Userinfo [id]`: Shows infos about an user;\n" +
			"`Info`: Shows infos about this bot;\n" +
			"**Staff** `Ban [id] [reason]`: Ban an user from the server [Optional reason];\n" +
			"**Staff** `Unban [id]`: Unban an user from the server;\n" +
			"**Staff** `Kick [id] [reason]`: Kick an user from the server [Optional reason];\n",
	}
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}