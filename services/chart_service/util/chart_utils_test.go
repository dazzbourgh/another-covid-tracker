package util

import (
	"another-covid-tracker.com/chart/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetChartData(t *testing.T) {
	input := types.Entries{
		types.Entry{Cases: 1, Vaccinations: 0.5, Date: TimeString("2021-January-02").Date()},
		types.Entry{Cases: 0, Vaccinations: 0, Date: TimeString("2021-January-01").Date()},
		types.Entry{Cases: 2, Vaccinations: 2, Date: TimeString("2021-January-03").Date()},
		types.Entry{Cases: 3, Vaccinations: 1.5, Date: TimeString("2021-January-04").Date()},
	}
	actual := GetChartData(input)
	assert.Equal(t, actual[0].Cases, float64(0))
	assert.Equal(t, input[0].Cases, float64(1))
}
