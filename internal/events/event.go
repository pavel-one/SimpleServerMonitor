package events

import "fmt"

type Interface interface {
	GetEvent() string
	GetData() any
}

type Chan chan Interface

func FormatEventName(topic string, eventName string) string {
	return fmt.Sprintf("event:%s:%s", topic, eventName)
}
