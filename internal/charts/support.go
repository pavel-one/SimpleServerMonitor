package charts

import "time"

func ModelsToChart(models []*Model, timeLayout string) *Chart {
	mapper := make(map[uint]*dataset)
	mapperTimes := make(map[time.Time]bool)
	var labels []string
	var datasets []*dataset

	for _, m := range models {
		t, err := time.Parse("2006-01-02 15:04:05", m.Time)
		if err != nil {
			continue
		}

		v, ok := mapper[m.SensorId]
		if ok {
			v.Data = append(v.Data, m.Temp)
		} else {
			mapper[m.SensorId] = &dataset{
				Label:    m.Name,
				SensorId: m.SensorId,
				Data:     []float32{m.Temp},
			}
		}

		_, ok = mapperTimes[t]
		if !ok {
			mapperTimes[t] = true
		}
	}

	for t := range mapperTimes {
		labels = append(labels, t.Format(timeLayout))
	}

	for _, ds := range mapper {
		datasets = append(datasets, ds)
	}

	chart := &Chart{
		Labels:   labels,
		Datasets: datasets,
	}

	return chart
}
