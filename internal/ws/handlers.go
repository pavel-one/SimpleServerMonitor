package ws

import (
	"github.com/goccy/go-json"
	"github.com/olahol/melody"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/sensors"
	"net/http"
)

func (s *Socket) SetDefault() {
	s.Http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := s.Server.HandleRequest(w, r); err != nil {
			s.Logger.Errorln(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	s.Server.HandleConnect(s.handleConnect)
	s.Server.HandleDisconnect(s.handleDisconnect)
	s.Server.HandleMessage(s.handleMessage)
}

func (s *Socket) handleConnect(sess *melody.Session) {
	s.Logger.Infoln("New Connection: ", sess.RemoteAddr())
	rep := sensors.NewSensorRepository(s.db)

	sens, err := rep.GetAll()
	if err != nil {
		s.Logger.Infoln("Error connection: ", err)
		if err := sess.Close(); err != nil {
			s.Logger.Errorln("Error session close: ", err)
			return
		}
		return
	}

	msg := events.NewTempAll(sens)
	b, err := json.Marshal(msg)
	if err != nil {
		s.Logger.Infoln("Error connection: ", err)
		if err := sess.Close(); err != nil {
			s.Logger.Errorln("Error session close: ", err)
			return
		}
		return
	}

	if err := sess.Write(b); err != nil {
		s.Logger.Infoln("Error connection: ", err)

		if err := sess.Close(); err != nil {
			s.Logger.Errorln("Error session close: ", err)
			return
		}
		return
	}

}

func (s *Socket) handleDisconnect(sess *melody.Session) {
	s.Logger.Infoln("Disconnection: ", sess.RemoteAddr())
}

func (s *Socket) handleMessage(sess *melody.Session, msg []byte) {
	s.Logger.Infoln("Message: ", sess.RemoteAddr(), string(msg))

	if err := s.Server.Broadcast([]byte("hi")); err != nil {
		s.Logger.Infoln("Broadcasting error: ", err)
	}
}
