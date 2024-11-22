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
	HelpCommandName = "help"

	helpPageCustomIDGroups = 2
)

var (
	HelpCustomID     = regexp.MustCompile(fmt.Sprintf("^/%s$", HelpCommandName))
	HelpPageCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/page$", HelpCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl
func getHelpSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     HelpCommandName,
		Description:              i18n.GetDefault("help.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("help.description"),
	}
}

func CraftHelpCustomID() string {
	return fmt.Sprintf("/%s", HelpCommandName)
}

func CraftHelpPageCustomID(commandName string) string {
	return fmt.Sprintf("/%s/%s/page", HelpCommandName, commandName)
}

func ExtractHelpPageCustomID(customID string) (string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, HelpPageCustomID,
		helpPageCustomIDGroups); ok {
		return groups[1], true
	}

	return "", false
}

func IsBelongsToHelp(customID string) bool {
	return regex.IsBelongTo(customID, HelpCustomID, HelpPageCustomID)
}
