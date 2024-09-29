package pubsub

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
)

type Observer struct {
	client  *redis.Client
	channel string
	queue   chan string
}

func NewObserver(client *redis.Client, channel string, queueSize int) *Observer {
	return &Observer{
		client:  client,
		channel: channel,
		queue:   make(chan string, queueSize),
	}
}

func (o *Observer) Subscribe(c context.Context) {
	subscriber := o.client.Subscribe(c, o.channel)
	defer subscriber.Close()

	for {
		select {
		case <-c.Done():
			return
		case msg := <-subscriber.Channel():
			if msg == nil {
				continue
			}
			log.Printf("chatch msg: %s", msg.String())
			o.queue <- msg.Payload
		}
	}
}

func (o *Observer) ProcessMessages(handler func(msg string)) {
	for msg := range o.queue {
		go handler(msg)
	}
}
