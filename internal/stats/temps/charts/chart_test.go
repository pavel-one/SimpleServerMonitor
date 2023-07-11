package charts

import (
	"errors"
	"fmt"
	"github.com/pavel-one/SimpleServerMonitor/tests"
	"testing"
)

func getRepository(time string) (*Repository, error) {
	db := tests.GetEmptyTestDB()
	if db == nil {
		return nil, errors.New("not getting db")
	}

	rep := NewRepository(db)

	q := fmt.Sprintf(`
BEGIN TRANSACTION;
INSERT INTO chips (id, name) VALUES (1, 'TestChip');
INSERT INTO sensors (id, name, high_temp, crit_temp, chip_id) VALUES (1, 'test', 80, 120, 1);
INSERT INTO sensors (id, name, high_temp, crit_temp, chip_id) VALUES (2, 'test1', 80, 120, 1);
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (20.00, 1, datetime('now', '-1 %[1]s'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (21.00, 1, datetime('now', '-2 %[1]s'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (22.00, 1, datetime('now', '-3 %[1]s'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (23.00, 1, datetime('now', '-4 %[1]s'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (24.00, 1, datetime('now', '-5 %[1]s'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (25.00, 2, datetime('now', '-1 %[1]s'));
END TRANSACTION;
`, time)
	_, err := db.Exec(q)
	if err != nil {
		return nil, err
	}

	return rep, nil
}

func checkChart(chart *Chart) error {
	if len(chart.Datasets) != 2 {
		return fmt.Errorf("Not correct dataset, need %d: %d", 2, len(chart.Datasets))
	}

	if len(chart.Datasets[0].Data) == 0 || len(chart.Datasets[1].Data) == 0 {
		return fmt.Errorf("Not correct: %d, %d", len(chart.Datasets[0].Data), len(chart.Datasets[1].Data))
	}

	return nil
}

func TestChartSensorsRepository_BySeconds(t *testing.T) {
	rep, err := getRepository("second")
	if err != nil {
		t.Fatal(err)
	}

	chart, err := rep.BySeconds()
	if err != nil {
		t.Fatal(err)
	}

	err = checkChart(chart)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChartSensorsRepository_ByMinutes(t *testing.T) {
	rep, err := getRepository("minute")
	if err != nil {
		t.Fatal(err)
	}

	chart, err := rep.ByMinutes()
	if err != nil {
		t.Fatal(err)
	}

	err = checkChart(chart)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChartSensorsRepository_ByHours(t *testing.T) {
	rep, err := getRepository("hour")
	if err != nil {
		t.Fatal(err)
	}

	chart, err := rep.ByHours()
	if err != nil {
		t.Fatal(err)
	}

	err = checkChart(chart)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChartSensorsRepository_ByDays(t *testing.T) {
	rep, err := getRepository("day")
	if err != nil {
		t.Fatal(err)
	}

	chart, err := rep.ByDays()
	if err != nil {
		t.Fatal(err)
	}

	err = checkChart(chart)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChartSensorsRepository_ByMonth(t *testing.T) {
	rep, err := getRepository("month")
	if err != nil {
		t.Fatal(err)
	}

	chart, err := rep.ByMonth()
	if err != nil {
		t.Fatal(err)
	}

	err = checkChart(chart)
	if err != nil {
		t.Fatal(err)
	}
}
