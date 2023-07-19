package temps

func duplicateSensor(key string, slice []*Stat) (find bool, index int) {
	for i, v := range slice {
		if v.Key == key {
			return true, i
		}
	}

	return false, 0
}
