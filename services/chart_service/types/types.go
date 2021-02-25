package types

import "time"

type Entry struct {
	IsoCode      string
	Cases        float32
	Vaccinations float32
	Date         time.Time
}

type Entries []Entry

func (entries *Entries) reduce() Entries {
	originalSize := len(*entries)
	resultArray := make(Entries, 10)
	step := originalSize / 10
	for i, j := 0, 0; i < originalSize; i, j = i+step, j+1 {
		resultArray[i] = (*entries)[j]
	}
	return resultArray
}
