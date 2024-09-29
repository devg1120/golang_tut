package pubsub

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
)

type Notifier struct {
	client *redis.Client
}

func NewNotifier(client *redis.Client) *Notifier {
	return &Notifier{client: client}
}

func (s *Notifier) PublishMessage(c context.Context, channel, message string) error {
	log.Print(channel)
	log.Print(message)
	return s.client.Publish(c, channel, message).Err()
}
