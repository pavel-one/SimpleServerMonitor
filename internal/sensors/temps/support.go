package temps

import "github.com/pavel-one/SimpleServerMonitor/internal/charts"

func duplicateSensor(key string, slice []*Stat) (find bool, index int) {
	for i, v := range slice {
		if v.Key == key {
			return true, i
		}
	}

	return false, 0
}

func duplicateCharts(key string, slice []charts.Data) (find bool, index int) {
	for i, v := range slice {
		if v.Name == key {
			return true, i
		}
	}

	return false, 0
}

func modelsToChart(models []Model) []charts.Data {
	var out []charts.Data

	for _, m := range models {
		find, index := duplicateCharts(m.Key, out)
		if find {
			out[index].Data = append(out[index].Data, []any{
				m.Time.Local().UnixMilli(),
				m.Temp,
			})
			continue
		}

		out = append(out, charts.Data{
			Name: m.Key,
			Data: [][]any{
				{
					m.Time.Local().UnixMilli(),
					m.Temp,
				},
			},
		})
	}

	return out
}
