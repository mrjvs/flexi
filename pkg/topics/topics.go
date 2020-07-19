package topics

const (
	MessageCreate            = "discord.event.guild.%s.channel.%s.message.create"
	MessageDelete            = "discord.event.guild.%s.channel.%s.message.%s.delete"
	MessageUpdate            = "discord.event.guild.%s.channel.%s.message.%s.update"
	MessageReactionAdd       = "discord.event.guild.%s.channel.%s.message.%s.reaction.add"
	MessageReactionRemove    = "discord.event.guild.%s.channel.%s.message.%s.reaction.remove"
	MessageReactionRemoveAll = "discord.event.guild.%s.channel.%s.message.%s.remove.all"
	MessageDeleteBulk        = "discord.event.guild.%s.channel.%s.message.delete.bulk"

	ChannelCreate         = "discord.event.guild.%s.channel.create"
	ChannelUpdate         = "discord.event.guild.%s.channel.%s.update"
	ChannelDelete         = "discord.event.guild.%s.channel.%s.delete"
	ChannelPinsUpdate     = "discord.event.guild.%s.channel.%s.pins.update"
	ChannelWebhooksUpdate = "discord.event.guild.%s.channel.%s.webhooks.update"

	GuildCreate    = "discord.event.guild.create"
	GuildUpdate    = "discord.event.guild.%s.update"
	GuildDelete    = "discord.event.guild.%s.delete"
	GuildBanAdd    = "discord.event.guild.%s.ban.add"
	GuildBanRemove = "discord.event.guild.%s.ban.remove"

	GuildMemberAdd    = "discord.event.guild.%s.member.add"
	GuildMemberUpdate = "discord.event.guild.%s.member.%s.update"
	GuildMemberRemove = "discord.event.guild.%s.member.%s.remove"

	GuildRoleCreate = "discord.event.guild.%s.role.create"
	GuildRoleUpdate = "discord.event.guild.%s.role.%s.update"
	GuildRoleDelete = "discord.event.guild.%s.role.%s.delete"

	GuildEmojisUpdate       = "discord.event.guild.%s.emojis.update"
	GuildIntegrationsUpdate = "discord.event.guild.%s.integrations.update"

	PresencesReplace = "discord.event.presences.replace"
	PresenceUpdate   = "discord.event.presence.update"

	// Probably unneeded
	RelationshipAdd         = "discord.event.relationship.add"
	RelationshipRemove      = "discord.event.relationship.remove"
	UserSettingsUpdate      = "discord.event.user.settings.update"
	UserGuildSettingsUpdate = "discord.event.user.guild.settings.update"
	UserNoteUpdate          = "discord.event.user.note.update"

	VoiceServerUpdate = "discord.event.guild.%s.voice.server.update"
	VoiceStateUpdate  = "discord.event.guild.%s.channel.%s.user.%s.voice.state.update"
)

const (
	GatewayMessageSend = "discord.message.create"
)

type SendMessage struct {
	GuildID   string
	ChannelID string
	Content   string
}
