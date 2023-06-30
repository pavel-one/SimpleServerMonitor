package ws

import (
	"github.com/olahol/melody"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
)

func worker(sess *melody.Session, ch events.Chan) {
	for true {
		if sess.IsClosed() {
			return
		}

	}
}
