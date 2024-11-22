package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	AboutCommandName = "about"
)

//nolint:nolintlint,exhaustive,lll,dupl
func getAboutSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     AboutCommandName,
		Description:              i18n.GetDefault("about.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("about.description"),
	}
}
