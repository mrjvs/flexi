package main

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mrjvs/flexi/pkg/topics"
)

func (g *gateway) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.MessageCreate, m.GuildID, m.ChannelID), msg)
}

func (g *gateway) onMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.MessageUpdate, m.GuildID, m.ChannelID, m.ID), msg)
}

func (g *gateway) onMessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.MessageDelete, m.GuildID, m.ChannelID, m.ID), msg)
}

func (g *gateway) onGuildCreate(s *discordgo.Session, m *discordgo.GuildCreate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildCreate), msg)
}

func (g *gateway) onGuildUpdate(s *discordgo.Session, m *discordgo.GuildUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildUpdate, m.ID), msg)
}

func (g *gateway) onGuildDelete(s *discordgo.Session, m *discordgo.GuildDelete) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildDelete, m.ID), msg)
}

func (g *gateway) onChannelCreate(s *discordgo.Session, m *discordgo.ChannelCreate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.ChannelCreate, m.GuildID), msg)
}

func (g *gateway) onChannelUpdate(s *discordgo.Session, m *discordgo.ChannelUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.ChannelUpdate, m.GuildID, m.ID), msg)
}

func (g *gateway) onChannelDelete(s *discordgo.Session, m *discordgo.ChannelDelete) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.ChannelDelete, m.GuildID, m.ID), msg)
}

func (g *gateway) onChannelPinsUpdate(s *discordgo.Session, m *discordgo.ChannelPinsUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.ChannelPinsUpdate, m.GuildID, m.ChannelID), msg)
}

func (g *gateway) onGuildBanAdd(s *discordgo.Session, m *discordgo.GuildBanAdd) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildBanAdd, m.GuildID), msg)
}

func (g *gateway) onGuildBanRemove(s *discordgo.Session, m *discordgo.GuildBanRemove) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildBanRemove, m.GuildID), msg)
}

func (g *gateway) onGuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildMemberAdd, m.GuildID), msg)
}

func (g *gateway) onGuildMemberUpdate(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildMemberUpdate, m.GuildID, m.User.ID), msg)
}

func (g *gateway) onGuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildMemberRemove, m.GuildID, m.User.ID), msg)
}

func (g *gateway) onGuildRoleCreate(s *discordgo.Session, m *discordgo.GuildRoleCreate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildRoleCreate, m.GuildID), msg)
}

func (g *gateway) onGuildRoleUpdate(s *discordgo.Session, m *discordgo.GuildRoleUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildRoleUpdate, m.GuildID, m.Role.ID), msg)
}

func (g *gateway) onGuildRoleDelete(s *discordgo.Session, m *discordgo.GuildRoleDelete) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildRoleDelete, m.GuildID, m.RoleID), msg)
}

func (g *gateway) onGuildEmojisUpdate(s *discordgo.Session, m *discordgo.GuildEmojisUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildEmojisUpdate, m.GuildID), msg)
}

func (g *gateway) onGuildIntegrationsUpdate(s *discordgo.Session, m *discordgo.GuildIntegrationsUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.GuildIntegrationsUpdate, m.GuildID), msg)
}

func (g *gateway) onMessageReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.MessageReactionAdd, m.GuildID, m.ChannelID, m.MessageID), msg)
}

func (g *gateway) onMessageReactionRemove(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.MessageReactionRemove, m.GuildID, m.ChannelID, m.MessageID), msg)
}

func (g *gateway) onMessageReactionRemoveAll(s *discordgo.Session, m *discordgo.MessageReactionRemoveAll) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.MessageReactionRemoveAll, m.GuildID, m.ChannelID, m.MessageID), msg)
}

func (g *gateway) onPresencesReplace(s *discordgo.Session, m *discordgo.PresencesReplace) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.PresencesReplace), msg)
}

func (g *gateway) onPresenceUpdate(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.PresenceUpdate), msg)
}

func (g *gateway) onRelationshipAdd(s *discordgo.Session, m *discordgo.RelationshipAdd) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.RelationshipAdd), msg)
}

func (g *gateway) onRelationshipRemove(s *discordgo.Session, m *discordgo.RelationshipRemove) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.RelationshipRemove), msg)
}

func (g *gateway) onUserSettingsUpdate(s *discordgo.Session, m *discordgo.UserSettingsUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.UserSettingsUpdate), msg)
}

func (g *gateway) onUserGuildSettingsUpdate(s *discordgo.Session, m *discordgo.UserGuildSettingsUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.UserGuildSettingsUpdate), msg)
}

func (g *gateway) onUserNoteUpdate(s *discordgo.Session, m *discordgo.UserNoteUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.UserNoteUpdate), msg)
}

func (g *gateway) onVoiceServerUpdate(s *discordgo.Session, m *discordgo.VoiceServerUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.VoiceServerUpdate, m.GuildID), msg)
}

func (g *gateway) onVoiceStateUpdate(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.VoiceStateUpdate, m.GuildID, m.ChannelID, m.UserID), msg)
}

func (g *gateway) onMessageDeleteBulk(s *discordgo.Session, m *discordgo.MessageDeleteBulk) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.MessageDeleteBulk, m.GuildID, m.ChannelID), msg)
}

func (g *gateway) onWebhooksUpdate(s *discordgo.Session, m *discordgo.WebhooksUpdate) {
	// handle message
	msg, _ := json.Marshal(m)
	g.nc.Publish(fmt.Sprintf(topics.ChannelWebhooksUpdate, m.GuildID, m.ChannelID), msg)
}
