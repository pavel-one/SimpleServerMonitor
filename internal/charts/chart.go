package charts

type Data struct {
	Name string  `json:"name"`
	Data [][]any `json:"data"`
}

type ChartRepository interface {
	GetBySeconds() ([]Data, error)
	GetByMinutes() ([]Data, error)
	GetByHours() ([]Data, error)
	GetByDays() ([]Data, error)
	GetByMonth() ([]Data, error)
}
