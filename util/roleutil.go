package util

import (
	"github.com/bwmarrin/discordgo"
)

func RoleGerarchyDifference(u1 *discordgo.Member, u2 *discordgo.Member, g *discordgo.Guild) int {
	u1MaxP, u2MaxP := -1, -1
	roleGerarchy := make(map[string]int)

	for _, rG := range g.Roles {
		roleGerarchy[rG.ID] = rG.Position
	}

	for _, r := range u1.Roles {
		p := roleGerarchy[r]
		if p > u1MaxP || u1MaxP == -1 {
			u1MaxP = p
		}
	}

	for _, r := range u2.Roles {
		p := roleGerarchy[r]
		if p > u2MaxP || u2MaxP == -1 {
			u2MaxP = p
		}
	}

	return u1MaxP - u2MaxP
}

func CanBan(u *discordgo.Member, g *discordgo.Guild) {

}