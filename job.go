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
	JobSlashCommandName  = "job"
	JobUserCommandName   = "Jobs"
	JobGetSubCommandName = "get"
	JobSetSubCommandName = "set"

	JobJobOptionName    = "job"
	JobLevelOptionName  = "level"
	JobServerOptionName = "server"

	JobMinLevel = 0
	JobMaxLevel = 200

	jobBookCustomIDGroups = 4
)

var (
	JobBookCustomID = regexp.MustCompile(fmt.
		Sprintf("^/%s/book\\?job=([a-z_]+)&server=([a-z_]+)&page=(\\d+)$", JobSlashCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl
func getJobSlashCommand() *discordgo.ApplicationCommand {
	var minLevel float64 = JobMinLevel
	return &discordgo.ApplicationCommand{
		Name:                     JobSlashCommandName,
		Description:              i18n.GetDefault("job.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("job.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     JobGetSubCommandName,
				Description:              i18n.GetDefault("job.get.description"),
				NameLocalizations:        *i18n.GetLocalizations("job.get.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("job.get.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     JobJobOptionName,
						Description:              i18n.GetDefault("job.get.job.description"),
						NameLocalizations:        *i18n.GetLocalizations("job.get.job.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.get.job.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:              JobServerOptionName,
						Description:       i18n.GetDefault("job.get.server.description"),
						NameLocalizations: *i18n.GetLocalizations("job.get.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.get.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     false,
						Autocomplete: true,
					},
				},
			},
			{
				Name:                     JobSetSubCommandName,
				Description:              i18n.GetDefault("job.set.description"),
				NameLocalizations:        *i18n.GetLocalizations("job.set.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("job.set.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     JobJobOptionName,
						Description:              i18n.GetDefault("job.set.job.description"),
						NameLocalizations:        *i18n.GetLocalizations("job.set.job.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.set.job.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     JobLevelOptionName,
						Description:              i18n.GetDefault("job.set.level.description"),
						NameLocalizations:        *i18n.GetLocalizations("job.set.level.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.set.level.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 true,
						MinValue:                 &minLevel,
						MaxValue:                 JobMaxLevel,
					},
					{
						Name:              JobServerOptionName,
						Description:       i18n.GetDefault("job.set.server.description"),
						NameLocalizations: *i18n.GetLocalizations("job.set.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.set.server.description",
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
func getJobUserCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     JobUserCommandName,
		Type:                     discordgo.UserApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
	}
}

func CraftJobBookCustomID(jobID, serverID string, page int) string {
	return fmt.Sprintf("/%s/book?job=%v&server=%v&page=%v",
		JobSlashCommandName, jobID, serverID, page)
}

func ExtractJobBookCustomID(customID string) (string, string, int, bool) {
	if groups, ok := regex.ExtractCustomID(customID, JobBookCustomID,
		jobBookCustomIDGroups); ok {
		jobID := groups[1]
		serverID := groups[2]
		page, errConv := strconv.Atoi(groups[3])
		if errConv != nil {
			return "", "", -1, false
		}

		return jobID, serverID, page, true
	}

	return "", "", -1, false
}

func IsBelongsToJob(customID string) bool {
	return regex.IsBelongTo(customID, JobBookCustomID)
}
