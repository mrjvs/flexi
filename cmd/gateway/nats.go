package main

import (
	"encoding/json"

	"github.com/mrjvs/flexi/pkg/topics"
	"github.com/nats-io/nats.go"
)

func (g *gateway) handleNatsMessage(msg *nats.Msg) {
	var p topics.SendMessage

	json.Unmarshal(msg.Data, &p)
	g.dg.ChannelMessageSend(p.ChannelID, p.Content)
}
