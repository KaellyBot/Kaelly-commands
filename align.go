package commands

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	"github.com/kaellybot/kaelly-commands/utils/regex"
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

	alignBookCustomIDGroups = 5

	alignAllValues = "_"
)

var (
	AlignBookCustomID = regexp.MustCompile(fmt.
		Sprintf("^/%s/book\\?city=([a-z_]+)&order=([a-z_]+)&server=([a-z_]+)&page=(\\d+)$", AlignSlashCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl
func getAlignSlashCommand() *discordgo.ApplicationCommand {
	var minLevel float64 = AlignmentMinLevel
	return &discordgo.ApplicationCommand{
		Name:                     AlignSlashCommandName,
		Description:              i18n.GetDefault("align.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("align.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     AlignGetSubCommandName,
				Description:              i18n.GetDefault("align.get.description"),
				NameLocalizations:        *i18n.GetLocalizations("align.get.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("align.get.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlignCityOptionName,
						Description:              i18n.GetDefault("align.get.city.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.get.city.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.city.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
					{
						Name:                     AlignOrderOptionName,
						Description:              i18n.GetDefault("align.get.order.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.get.order.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.order.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
					{
						Name:              AlignServerOptionName,
						Description:       i18n.GetDefault("align.get.server.description"),
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
				Description:              i18n.GetDefault("align.set.description"),
				NameLocalizations:        *i18n.GetLocalizations("align.set.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("align.set.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlignCityOptionName,
						Description:              i18n.GetDefault("align.set.city.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.set.city.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.city.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     AlignOrderOptionName,
						Description:              i18n.GetDefault("align.set.order.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.set.order.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.order.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     AlignLevelOptionName,
						Description:              i18n.GetDefault("align.set.level.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.set.level.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.level.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 true,
						MinValue:                 &minLevel,
						MaxValue:                 AlignmentMaxLevel,
					},
					{
						Name:              AlignServerOptionName,
						Description:       i18n.GetDefault("align.set.server.description"),
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

func CraftAlignBookCustomID(cityID, orderID, serverID string, page int) string {
	if len(cityID) == 0 {
		cityID = alignAllValues
	}

	if len(orderID) == 0 {
		orderID = alignAllValues
	}

	return fmt.Sprintf("/%s/book?city=%v&order=%v&server=%v&page=%v",
		AlignSlashCommandName, cityID, orderID, serverID, page)
}

func ExtractAlignBookCustomID(customID string,
) (string, string, string, int, bool) {
	if groups, ok := regex.ExtractCustomID(customID, AlignBookCustomID,
		alignBookCustomIDGroups); ok {
		cityID := groups[1]
		orderID := groups[2]
		serverID := groups[3]

		if cityID == alignAllValues {
			cityID = ""
		}

		if orderID == alignAllValues {
			orderID = ""
		}

		page, errConv := strconv.Atoi(groups[4])
		if errConv != nil {
			return "", "", "", -1, false
		}

		return cityID, orderID, serverID, page, true
	}

	return "", "", "", -1, false
}

func IsBelongsToAlign(customID string) bool {
	return regex.IsBelongTo(customID, AlignBookCustomID)
}
