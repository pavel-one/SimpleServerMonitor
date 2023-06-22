package sensors

import (
	"errors"
	"fmt"
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
	Data     []DataItem `json:"data" db:"-"`
}

type SensorRepository struct {
	DB *sqlx.DB
}

func NewSensorRepository(DB *sqlx.DB) *SensorRepository {
	return &SensorRepository{DB: DB}
}

// AddTemp create or find sensor and add temp
func (r *SensorRepository) AddTemp(sensor *Sensor) error {
	model, err := r.FindWithColumn("name", sensor.Name)

	if err != nil {
		model = &Model{
			Name:     sensor.Name,
			HighTemp: sensor.HighTemp,
			CritTemp: sensor.CritTemp,
		}

		model, err = r.Create(model)
		if err != nil {
			return err
		}
	}

	if model.Data == nil {
		model.Data = []DataItem{}
	}

	newData, err := r.addData(model.ID, sensor.Temp)
	if err != nil {
		return fmt.Errorf("error create data: %s", err)
	}

	model.Data = append(model.Data, newData)

	return nil
}

// Create save model to database
func (r *SensorRepository) Create(model *Model) (*Model, error) {
	_, err := r.DB.NamedExec("INSERT INTO sensors (name, high_temp, crit_temp) VALUES (:name, :high_temp, :crit_temp)", model)
	if err != nil {
		return nil, fmt.Errorf("error insert: %s", err)
	}

	updateModel, err := r.FindWithColumn("name", model.Name)
	if err != nil {
		return nil, fmt.Errorf("error update: %s", err)
	}

	return updateModel, nil
}

// Find get one model
func (r *SensorRepository) Find(id uint) (*Model, error) {
	model := Model{}

	if err := r.DB.Get(&model, "SELECT * FROM sensors WHERE id=$1 ORDER BY id DESC LIMIT 1", id); err != nil {
		return nil, err
	}

	if model.ID == 0 {
		return nil, errors.New("not found model")
	}

	model.Data = r.getData(model.ID)

	return &model, nil
}

// FindWithColumn find model from column name
func (r *SensorRepository) FindWithColumn(column string, value any) (*Model, error) {
	model := new(Model)

	q := fmt.Sprintf("SELECT * FROM sensors WHERE %s=$1 ORDER BY id DESC LIMIT 1", column)
	if err := r.DB.Get(model, q, value); err != nil {
		return nil, err
	}

	if model.ID == 0 {
		return nil, errors.New("not found")
	}

	model.Data = r.getData(model.ID)

	return model, nil
}

func (r *SensorRepository) GetAll() ([]*Model, error) {
	var models []*Model

	if err := r.DB.Select(&models, "SELECT * FROM sensors ORDER BY name"); err != nil {
		return nil, err
	}

	for _, item := range models {
		item.Data = r.getData(item.ID)
	}

	return models, nil
}

// TODO: move to another repository
func (r *SensorRepository) getData(sensorId uint) []DataItem {
	var data []DataItem

	err := r.DB.Select(&data, `
		SELECT temp, created_at FROM sensors_data 
            WHERE sensor_id=$1 
            ORDER BY created_at DESC`, sensorId)

	if err != nil {
		return nil
	}

	return data
}

// TODO: move to another repository
func (r *SensorRepository) addData(sensorId uint, temp float32) (DataItem, error) {
	item := DataItem{
		Temp:      temp,
		CreatedAt: time.Now(),
	}

	_, err := r.DB.Exec("INSERT INTO sensors_data (temp, sensor_id) VALUES ($1, $2)", item.Temp, sensorId)
	if err != nil {
		return item, err
	}

	return item, nil
}
