package ws

import (
	"github.com/olahol/melody"
	"net/http"
)

func (s *Socket) DefaultHandlers() {
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
