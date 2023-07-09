package charts

import "time"

func mapToChart(models []*Model) *Chart {
	chart := &Chart{
		DateStart: time.Now(),
		Datasets:  make([]*dataset, 0),
	}

	for _, m := range models {
		i, ok := findDataset(chart.Datasets, m.SensorID)

		if chart.DateStart.Sub(m.Time) > 0 {
			chart.DateStart = m.Time
		}

		if !ok {
			chart.Datasets = append(chart.Datasets, &dataset{
				Name:     m.Name,
				SensorID: m.SensorID,
				Data: [][]any{
					{m.Time.UnixMilli(), m.Temp},
				},
			})
			continue
		}

		chart.Datasets[i].Data = append(chart.Datasets[i].Data, []any{m.Time.UnixMilli(), m.Temp})
	}

	return chart
}

func findDataset(datasets []*dataset, sensorID uint) (index int, exists bool) {
	for i, v := range datasets {
		if v.SensorID == sensorID {
			return i, true
		}
	}

	return 0, false
}
