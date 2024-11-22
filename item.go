// nolint:nolintlint,dupl
package commands

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	"github.com/kaellybot/kaelly-commands/utils/regex"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	ItemCommandName = "item"

	ItemQueryOptionName = "query"

	itemCustomIDGroups        = 2
	itemEffectsCustomIDGroups = 3
	itemRecipeCustomIDGroups  = 3
)

var (
	ItemCustomID        = regexp.MustCompile(fmt.Sprintf("^/%s\\?type=(\\w+)$", ItemCommandName))
	ItemEffectsCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/effects\\?type=(\\w+)$", ItemCommandName))
	ItemRecipeCustomID  = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/recipe\\?type=(\\w+)$", ItemCommandName))
)

//nolint:exhaustive,lll,funlen
func getItemSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     ItemCommandName,
		Description:              i18n.GetDefault("item.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("item.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     ItemQueryOptionName,
				Description:              i18n.GetDefault("item.query.description"),
				NameLocalizations:        *i18n.GetLocalizations("item.query.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("item.query.description"),
				Type:                     discordgo.ApplicationCommandOptionString,
				Required:                 true,
				Autocomplete:             true,
			},
		},
	}
}

func CraftItemCustomID(itemType string) string {
	return fmt.Sprintf("/%s?type=%s", ItemCommandName, itemType)
}

func CraftItemEffectsCustomID(itemID, itemType string) string {
	return fmt.Sprintf("/%s/%s/effects?type=%s", ItemCommandName, itemID, itemType)
}

func CraftItemRecipeCustomID(itemID, itemType string) string {
	return fmt.Sprintf("/%s/%s/recipe?type=%s", ItemCommandName, itemID, itemType)
}

func ExtractItemCustomID(customID string) (string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, ItemCustomID,
		itemCustomIDGroups); ok {
		return groups[1], true
	}

	return "", false
}

func ExtractItemEffectsCustomID(customID string) (string, string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, ItemEffectsCustomID,
		itemEffectsCustomIDGroups); ok {
		return groups[1], groups[2], true
	}

	return "", "", false
}

func ExtractItemRecipeCustomID(customID string) (string, string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, ItemRecipeCustomID,
		itemRecipeCustomIDGroups); ok {
		return groups[1], groups[2], true
	}

	return "", "", false
}

func IsBelongsToItem(customID string) bool {
	return regex.IsBelongTo(customID, ItemCustomID,
		ItemEffectsCustomID, ItemRecipeCustomID)
}
