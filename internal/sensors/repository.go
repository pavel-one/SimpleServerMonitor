package sensors

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type DataItem struct {
	Temp      float32   `json:"temp" db:"temp"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Model struct {
	ID       uint       `json:"id" db:"id"`
	Name     string     `json:"name" db:"name"`
	HighTemp float32    `json:"high_temp" db:"high_temp"`
	CritTemp float32    `json:"crit_temp" db:"crit_temp"`
	Data     []DataItem `json:"data" db:"data"`
}

type SensorRepository struct {
	DB *sqlx.DB
}

func NewSensorRepository(DB *sqlx.DB) *SensorRepository {
	return &SensorRepository{DB: DB}
}

func (r *SensorRepository) Find(id uint) (*Model, error) {
	model := Model{}
	var dataTemps []DataItem

	if err := r.DB.Get(&model, "SELECT * FROM sensors WHERE id=$1 ORDER BY id DESC LIMIT 1", id); err != nil {
		return nil, err
	}

	if model.ID == 0 {
		return nil, errors.New("not found model")
	}

	err := r.DB.Select(&dataTemps, `
		SELECT temp, created_at FROM sensors_data 
            WHERE sensor_id=$1 
            ORDER BY created_at DESC`, model.ID)

	if err != nil {
		return nil, err
	}

	model.Data = dataTemps

	return &model, nil
}
