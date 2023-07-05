package sensors

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type SensorRepository struct {
	DB *sqlx.DB
}

func NewSensorRepository(DB *sqlx.DB) *SensorRepository {
	return &SensorRepository{DB: DB}
}

// AddTemp create or find sensor and add temp
func (r *SensorRepository) AddTemp(sensor *Sensor, chipName string) (*Model, error) {
	chipModel := new(ChipModel)

	if err := r.DB.Get(chipModel, "SELECT id, name FROM chips WHERE name=$1", chipName); err != nil {
		_, err = r.DB.Exec("INSERT INTO chips (name) VALUES ($1)", chipName)
		if err != nil {
			return nil, err
		}

		err = r.DB.Get(chipModel, "SELECT id, name FROM chips WHERE name=$1", chipName)
		if err != nil {
			return nil, err
		}
	}

	model, err := r.FindWithColumn("name", sensor.Name)

	if err != nil {
		model = &Model{
			Name:     sensor.Name,
			HighTemp: sensor.HighTemp,
			CritTemp: sensor.CritTemp,
			ChipId:   chipModel.ID,
		}

		model, err = r.Create(model)
		if err != nil {
			return nil, err
		}
	}

	if model.Data == nil {
		model.Data = []DataItem{}
	}

	newData, err := r.addData(model.ID, sensor.Temp)
	if err != nil {
		return nil, fmt.Errorf("error create data: %s", err)
	}

	model.Data = append(model.Data, newData)

	return model, nil
}

// Create save model to database
func (r *SensorRepository) Create(model *Model) (*Model, error) {
	_, err := r.DB.NamedExec("INSERT INTO sensors (name, high_temp, crit_temp, chip_id) VALUES (:name, :high_temp, :crit_temp, :chip_id)", model)
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
