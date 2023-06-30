package workers

import (
	"github.com/olahol/melody"
	"github.com/pavel-one/SimpleServerMonitor/internal/Logger"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/sensors"
	"github.com/pavel-one/SimpleServerMonitor/internal/sql"
)

func WebsocketWorker(sess melody.Session, ch events.Chan) error {
	log := Logger.NewLogger("WsWorker")
	db, err := sql.Connect("db")
	if err != nil {
		return err
	}
	rep := sensors.NewSensorRepository(db)
	log.Infoln("Run user ws pulling", sess.RemoteAddr())

	for i := range ch {

	}
}
