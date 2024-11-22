package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	PosCommandName         = "pos"
	PosDimensionOptionName = "dimension"
	PosServerOptionName    = "server"
)

//nolint:nolintlint,exhaustive,lll,dupl
func getPosSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     PosCommandName,
		Description:              i18n.GetDefault("pos.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("pos.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     PosDimensionOptionName,
				Description:              i18n.GetDefault("pos.dimension.description"),
				NameLocalizations:        *i18n.GetLocalizations("pos.dimension.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("pos.dimension.description"),
				Type:                     discordgo.ApplicationCommandOptionString,
				Required:                 false,
				Autocomplete:             true,
			},
			{
				Name:              PosServerOptionName,
				Description:       i18n.GetDefault("pos.server.description"),
				NameLocalizations: *i18n.GetLocalizations("pos.server.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("pos.server.description",
					i18n.Vars{"game": constants.GetGame()}),
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     false,
				Autocomplete: true,
			},
		},
	}
}
