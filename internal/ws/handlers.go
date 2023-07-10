package ws

import (
	"github.com/goccy/go-json"
	"github.com/olahol/melody"
	"github.com/pavel-one/SimpleServerMonitor/internal/charts"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"net/http"
)

// SetDefault handlers methods
func (s *Socket) SetDefault() {
	s.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	rep := charts.NewRepository(s.db)

	chart, err := rep.BySeconds()
	if err != nil {
		s.Logger.Infoln("Error connection: ", err)
		if err := sess.Close(); err != nil {
			s.Logger.Errorln("Error session close: ", err)
			return
		}
		return
	}

	msg := events.NewChart(chart, events.LoadTempEvent)
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
	e := new(events.Event)

	if err := json.Unmarshal(msg, e); err != nil {
		s.Logger.Errorln("Message is not is event", sess.RemoteAddr(), string(msg))
		return
	}

	if e.Channel == events.TempTopic && e.Event == events.LoadTempEvent {
		typ, ok := e.Data.(string)
		if !ok {
			s.Logger.Errorln("Message is not is getting correct data", sess.RemoteAddr(), string(msg))
			return
		}
		rep := charts.NewRepository(s.db)

		var chart *charts.Chart
		var err error
		switch typ {
		case charts.TypeSecond:
			chart, err = rep.BySeconds()
		case charts.TypeMinute:
			chart, err = rep.ByMinutes()
		case charts.TypeHour:
			chart, err = rep.ByHours()
		case charts.TypeDay:
			chart, err = rep.ByDays()
		case charts.TypeMonth:
			chart, err = rep.ByMonth()
		default:
			chart, err = rep.BySeconds()
		}

		if err != nil {
			s.Logger.Errorln("Failed getting data:", err)
			return
		}

		msg := events.NewChart(chart, events.LoadTempEvent)
		b, err := json.Marshal(msg)
		if err != nil {
			s.Logger.Errorln("Failed getting data:", err)
			return
		}

		if err := sess.Write(b); err != nil {
			s.Logger.Errorln("Failed send data:", err)
			return
		}
	}
}
