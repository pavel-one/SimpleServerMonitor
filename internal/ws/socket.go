package ws

import (
	"github.com/olahol/melody"
	"github.com/pavel-one/sensors/internal/Logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Socket struct {
	Server *melody.Melody
	Http   *http.ServeMux
	Port   int
	Logger *zap.SugaredLogger
}

func NewServer(port int, serverName string) *Socket {
	m := melody.New()
	h := http.NewServeMux()
	logger := Logger.NewLogger(serverName)

	return &Socket{
		Server: m,
		Http:   h,
		Port:   port,
		Logger: logger,
	}
}

func (s *Socket) Run() error {
	s.Logger.Infof("Server running on port %d", s.Port)

	if err := http.ListenAndServe(":"+strconv.Itoa(s.Port), s.Http); err != nil {
		return err
	}

	return nil
}
