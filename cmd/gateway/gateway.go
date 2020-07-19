package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nats-io/nats.go"
)

type gateway struct {
	dg *discordgo.Session
	nc *nats.Conn
}
