package memory

import "github.com/shirou/gopsutil/v3/mem"

type Stat struct {
	Total       uint64
	UsedPercent float64
	Used        uint64
	Free        uint64
}

type StatWrapper struct {
	Memory *Stat
	Swap   *Stat
}

func GetStats() (*StatWrapper, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	swap, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}

	return &StatWrapper{
		Memory: &Stat{
			Total:       memory.Total,
			UsedPercent: memory.UsedPercent,
			Used:        memory.Used,
			Free:        memory.Free,
		},
		Swap: &Stat{
			Total:       swap.Total,
			UsedPercent: swap.UsedPercent,
			Used:        swap.Used,
			Free:        swap.Free,
		},
	}, nil
}
