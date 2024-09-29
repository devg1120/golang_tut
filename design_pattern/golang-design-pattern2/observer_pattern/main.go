package main

import (
	"encoding/json"
	"github.com/cocoide/golang-design-pattern/observer_pattern/entity"
	o "github.com/cocoide/golang-design-pattern/observer_pattern/observer"
	"github.com/cocoide/golang-design-pattern/observer_pattern/observer/pubsub"
	"log"
	"time"
)

const (
	UserEventChannel = "user_event_channel"
)

func main() {
	redis := pubsub.NewClient()
	type AchieveEvent struct {
		Achievement entity.Achievement
		UpdatedAt   time.Time
	}
	observer := o.NewUserObserver[AchieveEvent](redis, UserEventChannel)
	notifier := o.NewUserNotifier[AchieveEvent](UserEventChannel)

	handler := func(msg string) {
		var event AchieveEvent
		if err := json.Unmarshal([]byte(msg), &event); err != nil {
			// logging
		}
		log.Printf("Acheive: %s", event.Achievement.String())
	}
	go observer.OnNotify(handler)

	notifier.Register(observer)
	event := AchieveEvent{Achievement: entity.LoginAchieve}
	notifier.Notify(event)

	select {}
}
