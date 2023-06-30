package messagess

import "github.com/pavel-one/SimpleServerMonitor/internal/sensors"

type Connection struct {
	Message
	Sensors []*sensors.Model
}

func NewConnection(sens []*sensors.Model) *Connection {
	msg := new(Connection)
	msg.Event = "load:all"
	msg.Sensors = sens

	return msg
}
