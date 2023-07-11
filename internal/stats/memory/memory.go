package memory

import (
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/SimpleServerMonitor/internal/stats/memory/sql"
	"github.com/shirou/gopsutil/v3/mem"
	"sync"
	"time"
)

type Service struct {
	actual *Model
	mu     sync.Mutex
	db     *sqlx.DB
}

func NewService() (*Service, error) {
	db, err := sql.Connect("db")
	if err != nil {
		return nil, err
	}

	s := new(Service)
	s.db = db

	if err := s.Update(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) Update() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	m, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	s.actual = &Model{
		Percent:   m.UsedPercent,
		Free:      m.Free,
		Total:     m.Total,
		CreatedAt: time.Now(),
	}
	return nil
}
