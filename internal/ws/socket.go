package ws

import (
	"github.com/jmoiron/sqlx"
	"github.com/olahol/melody"
	"github.com/pavel-one/SimpleServerMonitor/internal/db"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Socket server wrapper
type Socket struct {
	Server *melody.Melody
	Mux    *http.ServeMux
	Port   int
	Logger *zap.SugaredLogger
	db     *sqlx.DB
	events events.Chan
}

// NewServer create new socket server
func NewServer(port int, serverName string, ch events.Chan) *Socket {
	m := melody.New()
	h := http.NewServeMux()

	return &Socket{
		Server: m,
		Mux:    h,
		Port:   port,
		Logger: logger.NewLogger(serverName),
		events: ch,
	}
}

// Run socket server
func (s *Socket) Run() error {
	d, err := db.DefaultConnection()
	if err != nil {
		return err
	}
	s.db = d

	s.Logger.Infof("Server running on port %d", s.Port)

	if err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(s.Port), s.Mux); err != nil {
		return err
	}

	return nil
}
