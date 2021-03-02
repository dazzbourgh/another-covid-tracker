package util

import (
	"another-covid-tracker.com/chart/types"
	"sort"
)

type ByDate types.Entries

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Less(i, j int) bool {
	return a[i].Date.Before(a[j].Date)
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func GetChartData(entries types.Entries) types.Entries {
	sorted := make(types.Entries, 0)
	for i := range entries {
		if entries[i].Vaccinations > 0 {
			sorted = append(sorted, entries[i])
		}
	}
	sort.Sort(ByDate(sorted))
	return sorted
}
