package ws

import (
	"github.com/jmoiron/sqlx"
	"github.com/olahol/melody"
	"github.com/pavel-one/SimpleServerMonitor/internal/Logger"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/sql"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Socket struct {
	Server *melody.Melody
	Http   *http.ServeMux
	Port   int
	Logger *zap.SugaredLogger
	db     *sqlx.DB
	events events.Chan
}

func NewServer(port int, serverName string, ch events.Chan) *Socket {
	m := melody.New()
	h := http.NewServeMux()
	logger := Logger.NewLogger(serverName)

	return &Socket{
		Server: m,
		Http:   h,
		Port:   port,
		Logger: logger,
		events: ch,
	}
}

func (s *Socket) Run() error {
	db, err := sql.Connect("db")
	if err != nil {
		return err
	}
	s.db = db

	s.Logger.Infof("Server running on port %d", s.Port)

	if err := http.ListenAndServe(":"+strconv.Itoa(s.Port), s.Http); err != nil {
		return err
	}

	return nil
}
