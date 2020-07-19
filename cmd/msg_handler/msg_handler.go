package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mrjvs/flexi/pkg/topics"
	"github.com/nats-io/nats.go"
)

type msgHandler struct {
	nc *nats.Conn
}

func main() {
	nc, err := nats.Connect("nats://localhost:4222", func(o *nats.Options) error {
		o.Name = "gateway"
		return nil
	})

	var h msgHandler
	h.nc = nc

	if err != nil {
		log.Fatal(err)
	}
	nc.Subscribe(fmt.Sprintf(topics.MessageCreate, "*", "*"), h.handleMsg)
	log.Println("running")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

func (h *msgHandler) handleMsg(msg *nats.Msg) {
	var m discordgo.Message
	err := json.Unmarshal(msg.Data, &m)
	if err != nil {
		log.Printf("Could not unmarshal data (%v)\n", err)
	}
	if strings.HasPrefix(m.Content, ";ping") {
		p := topics.SendMessage{
			GuildID:   m.GuildID,
			ChannelID: m.ChannelID,
			Content:   "pong",
		}

		payload, _ := json.Marshal(p)

		h.nc.Publish(topics.GatewayMessageSend, payload)
	}
}
