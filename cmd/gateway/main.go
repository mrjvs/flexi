package main

//go:generate go run generated/gen.go

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mrjvs/flexi/pkg/topics"
	"github.com/nats-io/nats.go"
)

var botToken = os.Getenv("BOT_TOKEN")

func main() {
	if botToken == "" {
		log.Fatalln("no bot token present.")
	}

	var g gateway

	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}
	defer dg.Close()
	g.dg = dg

	nc, err := nats.Connect("nats://localhost:4222", func(o *nats.Options) error {
		o.Name = "gateway"
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	g.nc = nc

	dg.AddHandler(g.onMessageCreate)
	dg.AddHandler(g.onMessageUpdate)
	dg.AddHandler(g.onMessageDelete)

	dg.AddHandler(g.onGuildCreate)
	dg.AddHandler(g.onGuildUpdate)
	dg.AddHandler(g.onGuildDelete)

	dg.AddHandler(g.onChannelCreate)
	dg.AddHandler(g.onChannelUpdate)
	dg.AddHandler(g.onChannelDelete)
	dg.AddHandler(g.onChannelPinsUpdate)

	dg.AddHandler(g.onGuildBanAdd)
	dg.AddHandler(g.onGuildBanRemove)
	dg.AddHandler(g.onGuildMemberAdd)
	dg.AddHandler(g.onGuildMemberUpdate)
	dg.AddHandler(g.onGuildMemberRemove)
	dg.AddHandler(g.onGuildRoleCreate)
	dg.AddHandler(g.onGuildRoleUpdate)
	dg.AddHandler(g.onGuildRoleDelete)
	dg.AddHandler(g.onGuildEmojisUpdate)
	dg.AddHandler(g.onGuildIntegrationsUpdate)

	dg.AddHandler(g.onMessageReactionAdd)
	dg.AddHandler(g.onMessageReactionRemove)
	dg.AddHandler(g.onMessageReactionRemoveAll)

	dg.AddHandler(g.onPresencesReplace)
	dg.AddHandler(g.onPresenceUpdate)

	dg.AddHandler(g.onRelationshipAdd)
	dg.AddHandler(g.onRelationshipRemove)

	dg.AddHandler(g.onUserSettingsUpdate)
	dg.AddHandler(g.onUserGuildSettingsUpdate)
	dg.AddHandler(g.onUserNoteUpdate)

	dg.AddHandler(g.onVoiceServerUpdate)
	dg.AddHandler(g.onVoiceStateUpdate)

	dg.AddHandler(g.onMessageDeleteBulk)
	dg.AddHandler(g.onWebhooksUpdate)

	nc.QueueSubscribe(topics.GatewayMessageSend, "gateway", g.handleNatsMessage)

	log.Println("running")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
