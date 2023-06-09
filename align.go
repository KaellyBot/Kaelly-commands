package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	AlignSlashCommandName  = "align"
	AlignUserCommandName   = "Alignments"
	AlignGetSubCommandName = "get"
	AlignSetSubCommandName = "set"

	AlignCityOptionName   = "city"
	AlignOrderOptionName  = "order"
	AlignLevelOptionName  = "level"
	AlignServerOptionName = "server"

	AlignmentMinLevel = 0
	AlignmentMaxLevel = 100
)

//nolint:nolintlint,exhaustive,lll,dupl
func getAlignSlashCommand() *discordgo.ApplicationCommand {
	var minLevel float64 = AlignmentMinLevel
	return &discordgo.ApplicationCommand{
		Name:                     AlignSlashCommandName,
		Description:              "align.description",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("align.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     AlignGetSubCommandName,
				Description:              "align.get.description",
				NameLocalizations:        *i18n.GetLocalizations("align.get.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("align.get.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlignCityOptionName,
						Description:              "align.get.city.description",
						NameLocalizations:        *i18n.GetLocalizations("align.get.city.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.city.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
					{
						Name:                     AlignOrderOptionName,
						Description:              "align.get.order.description",
						NameLocalizations:        *i18n.GetLocalizations("align.get.order.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.order.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
					{
						Name:              AlignServerOptionName,
						Description:       "align.get.server.description",
						NameLocalizations: *i18n.GetLocalizations("align.get.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     false,
						Autocomplete: true,
					},
				},
			},
			{
				Name:                     AlignSetSubCommandName,
				Description:              "align.set.description",
				NameLocalizations:        *i18n.GetLocalizations("align.set.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("align.set.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlignCityOptionName,
						Description:              "align.set.city.description",
						NameLocalizations:        *i18n.GetLocalizations("align.set.city.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.city.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     AlignOrderOptionName,
						Description:              "align.set.order.description",
						NameLocalizations:        *i18n.GetLocalizations("align.set.order.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.order.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     AlignLevelOptionName,
						Description:              "align.set.level.description",
						NameLocalizations:        *i18n.GetLocalizations("align.set.level.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.level.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 true,
						MinValue:                 &minLevel,
						MaxValue:                 AlignmentMaxLevel,
					},
					{
						Name:              AlignServerOptionName,
						Description:       "align.set.server.description",
						NameLocalizations: *i18n.GetLocalizations("align.set.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     false,
						Autocomplete: true,
					},
				},
			},
		},
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getAlignUserCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     AlignUserCommandName,
		Type:                     discordgo.UserApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
	}
}
