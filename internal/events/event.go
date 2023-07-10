package events

import "fmt"

type Event struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Data    any    `json:"data"`
}

// Interface for event
type Interface interface {
	GetEvent() string
	GetData() any
}

// Chan channel for events
type Chan chan Interface

// FormatEventName create event name from topic and eventName
func FormatEventName(topic string, eventName string) string {
	return fmt.Sprintf("event:%s:%s", topic, eventName)
}
