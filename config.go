package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	ConfigCommandName           = "config"
	ConfigGetSubCommandName     = "get"
	ConfigAlmanaxSubCommandName = "almanax"
	ConfigRSSSubCommandName     = "rss"
	ConfigServerSubCommandName  = "server"
	ConfigTwitchSubCommandName  = "twitch"
	ConfigTwitterSubCommandName = "twitter"
	ConfigYoutubeSubCommandName = "youtube"

	ConfigServerOptionName         = "server"
	ConfigChannelOptionName        = "channel"
	ConfigFeedTypeOptionName       = "type"
	ConfigStreamerOptionName       = "streamer"
	ConfigTwitterAccountOptionName = "twitterAccount"
	ConfigVideastOptionName        = "videast"
	ConfigLanguageOptionName       = "language"
	ConfigEnabledOptionName        = "enabled"
)

//nolint:nolintlint,exhaustive,lll,dupl,funlen
func getConfigSlashCommand(localChoices []*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     ConfigCommandName,
		Description:              "config.description",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetManageServerPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("config.description",
			i18n.Vars{"game": constants.GetGame()}),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     ConfigGetSubCommandName,
				Description:              "config.get.description",
				NameLocalizations:        *i18n.GetLocalizations("config.get.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.get.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:                     ConfigAlmanaxSubCommandName,
				Description:              "config.almanax.description",
				NameLocalizations:        *i18n.GetLocalizations("config.almanax.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.almanax.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              "config.almanax.enabled.description",
						NameLocalizations:        *i18n.GetLocalizations("config.almanax.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.almanax.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigLanguageOptionName,
						Description:              "config.almanax.language.description",
						NameLocalizations:        *i18n.GetLocalizations("config.almanax.language.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.almanax.language.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 false,
						Choices:                  localChoices,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              "config.almanax.channel.description",
						NameLocalizations:        *i18n.GetLocalizations("config.almanax.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.almanax.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
			{
				Name:                     ConfigRSSSubCommandName,
				Description:              "config.rss.description",
				NameLocalizations:        *i18n.GetLocalizations("config.rss.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.rss.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              "config.rss.enabled.description",
						NameLocalizations:        *i18n.GetLocalizations("config.rss.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.rss.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigFeedTypeOptionName,
						Description:              "config.rss.feedtype.description",
						NameLocalizations:        *i18n.GetLocalizations("config.rss.feedtype.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.rss.feedtype.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     ConfigLanguageOptionName,
						Description:              "config.rss.language.description",
						NameLocalizations:        *i18n.GetLocalizations("config.rss.language.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.rss.language.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 false,
						Choices:                  localChoices,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              "config.rss.channel.description",
						NameLocalizations:        *i18n.GetLocalizations("config.rss.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.rss.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
			{
				Name:              ConfigServerSubCommandName,
				Description:       "config.server.description",
				NameLocalizations: *i18n.GetLocalizations("config.server.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.server.description",
					i18n.Vars{"game": constants.GetGame()}),
				Type: discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:              ConfigServerOptionName,
						Description:       "config.server.server.description",
						NameLocalizations: *i18n.GetLocalizations("config.server.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.server.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     true,
						Autocomplete: true,
					},
					{
						Name:              ConfigChannelOptionName,
						Description:       "config.server.channel.description",
						NameLocalizations: *i18n.GetLocalizations("config.server.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.server.channel.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:     discordgo.ApplicationCommandOptionChannel,
						Required: false,
					},
				},
			},
			{
				Name:                     ConfigTwitchSubCommandName,
				Description:              "config.twitch.description",
				NameLocalizations:        *i18n.GetLocalizations("config.twitch.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.twitch.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              "config.twitch.enabled.description",
						NameLocalizations:        *i18n.GetLocalizations("config.twitch.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitch.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigStreamerOptionName,
						Description:              "config.twitch.streamer.description",
						NameLocalizations:        *i18n.GetLocalizations("config.twitch.streamer.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitch.streamer.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              "config.twitch.channel.description",
						NameLocalizations:        *i18n.GetLocalizations("config.twitch.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitch.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
			{
				Name:              ConfigTwitterSubCommandName,
				Description:       "config.twitter.description",
				NameLocalizations: *i18n.GetLocalizations("config.twitter.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.description",
					i18n.Vars{"game": constants.GetGame()}),
				Type: discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              "config.twitter.enabled.description",
						NameLocalizations:        *i18n.GetLocalizations("config.twitter.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigTwitterAccountOptionName,
						Description:              "config.twitter.account.description",
						NameLocalizations:        *i18n.GetLocalizations("config.twitter.account.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.account.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              "config.twitter.channel.description",
						NameLocalizations:        *i18n.GetLocalizations("config.twitter.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
			{
				Name:                     ConfigYoutubeSubCommandName,
				Description:              "config.youtube.description",
				NameLocalizations:        *i18n.GetLocalizations("config.youtube.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.youtube.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              "config.youtube.enabled.description",
						NameLocalizations:        *i18n.GetLocalizations("config.youtube.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.youtube.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigVideastOptionName,
						Description:              "config.youtube.videast.description",
						NameLocalizations:        *i18n.GetLocalizations("config.youtube.videast.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.youtube.videast.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              "config.youtube.channel.description",
						NameLocalizations:        *i18n.GetLocalizations("config.youtube.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.youtube.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
		},
	}
}
