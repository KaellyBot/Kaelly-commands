package commands

import "github.com/bwmarrin/discordgo"

type GetLocalChoicesFn func() []*discordgo.ApplicationCommandOptionChoice

func GetCommands(localChoices []*discordgo.ApplicationCommandOptionChoice) []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		getAboutSlashCommand(),
		getAlignSlashCommand(),
		getAlignUserCommand(),
		getAlmanaxSlashCommand(),
		getConfigSlashCommand(localChoices),
		getItemSlashCommand(),
		getJobSlashCommand(),
		getJobUserCommand(),
		getPosSlashCommand(),
		getSetSlashCommand(),
	}
}
