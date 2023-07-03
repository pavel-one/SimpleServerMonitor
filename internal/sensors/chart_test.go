package sensors

import (
	"github.com/pavel-one/SimpleServerMonitor/tests"
	"testing"
)

func TestChartSensorsRepository_BySeconds(t *testing.T) {
	db := tests.GetEmptyTestDB()
	if db == nil {
		t.Fatal("not init database")
	}

	chartRep := NewChartSensorsRepository(db)

	_, err := db.Exec(`
BEGIN TRANSACTION;
INSERT INTO sensors (id, name, high_temp, crit_temp) VALUES (1, 'test', 80, 120);
INSERT INTO sensors (id, name, high_temp, crit_temp) VALUES (2, 'test1', 80, 120);
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (20.00, 1, datetime('now', '-5 second'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (21.00, 1, datetime('now', '-10 second'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (22.00, 1, datetime('now', '-15 second'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (23.00, 1, datetime('now', '-20 second'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (24.00, 1, datetime('now', '-25 second'));
INSERT INTO sensors_data (temp, sensor_id, created_at) VALUES (25.00, 2, datetime('now', '-5 second'));
END TRANSACTION;
`)
	if err != nil {
		t.Fatal(err)
	}

	chartRep.BySeconds()

}
