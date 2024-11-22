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
	MapCommandName = "map"

	MapNumberOptionName = "number"

	MapNumberMin = 1
	MapNumberMax = 50

	mapNormalCustomIDGroups   = 2
	mapTacticalCustomIDGroups = 2
)

var (
	MapNormalCustomID   = regexp.MustCompile(fmt.Sprintf("^/%s/(\\d+)\\?type=normal$", MapCommandName))
	MapTacticalCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\d+)\\?type=tactical$", MapCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl
func getMapSlashCommand() *discordgo.ApplicationCommand {
	var mapNumberMin float64 = MapNumberMin
	return &discordgo.ApplicationCommand{
		Name:                     MapCommandName,
		Description:              i18n.GetDefault("map.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("map.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     MapNumberOptionName,
				Description:              i18n.GetDefault("map.number.description"),
				NameLocalizations:        *i18n.GetLocalizations("map.number.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("map.number.description"),
				Type:                     discordgo.ApplicationCommandOptionInteger,
				Required:                 false,
				MinValue:                 &mapNumberMin,
				MaxValue:                 MapNumberMax,
			},
		},
	}
}

func CraftMapNormalCustomID(mapNumber int64) string {
	return fmt.Sprintf("/%s/%v?type=normal", MapCommandName, mapNumber)
}

func CraftMapTacticalCustomID(mapNumber int64) string {
	return fmt.Sprintf("/%s/%v?type=tactical", MapCommandName, mapNumber)
}

func ExtractMapNormalCustomID(customID string) (int64, bool) {
	if groups, ok := regex.ExtractCustomID(customID, MapNormalCustomID,
		mapNormalCustomIDGroups); ok {
		mapNumber, err := strconv.Atoi(groups[1])
		if err != nil {
			return -1, false
		}
		return int64(mapNumber), true
	}

	return -1, false
}

func ExtractMapTacticalCustomID(customID string) (int64, bool) {
	if groups, ok := regex.ExtractCustomID(customID, MapTacticalCustomID,
		mapTacticalCustomIDGroups); ok {
		mapNumber, err := strconv.Atoi(groups[1])
		if err != nil {
			return -1, false
		}
		return int64(mapNumber), true
	}

	return -1, false
}

func IsBelongsToMap(customID string) bool {
	return regex.IsBelongTo(customID, MapNormalCustomID, MapTacticalCustomID)
}
