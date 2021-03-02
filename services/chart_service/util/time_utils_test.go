package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateString(t *testing.T) {
	date := MyTime(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC))
	dateString := date.DateString()
	assert.Equal(t, "2021-01-01", dateString)
}

func TestDate(t *testing.T) {
	dateString := TimeString("2021-January-01")
	date := dateString.Date()
	assert.Equal(t, time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), date)
}
