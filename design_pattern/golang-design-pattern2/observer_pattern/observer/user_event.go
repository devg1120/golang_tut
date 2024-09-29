package observer

import (
	"encoding/json"
	"github.com/cocoide/golang-design-pattern/observer_pattern/observer/pubsub"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
)

type userObserver[event any] struct {
	pubsub  *pubsub.Observer
	channel string
}

func NewUserObserver[event any](client *redis.Client, channel string) Observer[event] {
	p := pubsub.NewObserver(client, channel, 10)
	return &userObserver[event]{pubsub: p}
}

func (i *userObserver[event]) OnNotify(handler func(msg string)) {
	i.pubsub.Subscribe(context.Background())
	i.pubsub.ProcessMessages(handler)
}

type userNotifier[event any] struct {
	ctx       context.Context
	channel   string // unique
	observers map[string]Observer[event]
	pubsub    *pubsub.Notifier
}

func NewUserNotifier[event any](channel string) Notifier[event] {
	return &userNotifier[event]{
		observers: make(map[string]Observer[event]),
		channel:   channel,
	}
}

func (i *userNotifier[event]) Register(o Observer[event]) {
	i.observers[i.channel] = o
}

func (i *userNotifier[event]) Unregister(o Observer[event]) {
	delete(i.observers, i.channel)
}

func (n *userNotifier[event]) Notify(e event) {
	for channel := range n.observers {
		msg, err := n.parseMsg(e)
		if err != nil {
			log.Print(err)
		}
		if err := n.pubsub.PublishMessage(n.ctx, channel, msg); err != nil {
			log.Print(err)
		}
	}
}

func (n *userNotifier[event]) parseMsg(e event) (string, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
