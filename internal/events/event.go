package events

type Interface interface {
	GetEvent() string
}

type event struct {
	Event string `json:"event"`
}

func (e *event) GetEvent() string {
	return e.Event
}

type Chan chan Interface
