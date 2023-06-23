package ws

import (
	"github.com/olahol/melody"
	"github.com/pavel-one/sensors/internal/events"
)

func worker(sess *melody.Session, ch events.Chan) {
	for true {
		if sess.IsClosed() {
			return
		}

	}
}
