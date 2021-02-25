package util

import (
	"another-covid-tracker.com/chart/types"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchEntries(t *testing.T) {
	now := time.Now()
	then := now.Add(48 * time.Hour)
	entries := make(types.Entries, 0)
	entries = append(entries, types.Entry{
		IsoCode:      "USA",
		Cases:        20,
		Vaccinations: 10,
		Date:         now,
	}, types.Entry{
		IsoCode:      "USA",
		Cases:        30,
		Vaccinations: 20,
		Date:         then,
	})

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(entries[0:1])
		w.Write(j)
	}))

	testServer.URL = testServer.URL + fmt.Sprintf("/USA?from=%s&to=%s", MyTime(now).DateString(), MyTime(then).DateString())
	url := testServer.URL
	defer testServer.Close()

	from := now
	to := now.Add(24 * time.Hour)

	result := FetchEntries(url)("USA", from, to)
	entry := result[0]
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "USA", entry.IsoCode)
	assert.Equal(t, float32(20), entry.Cases)
	assert.Equal(t, float32(10), entry.Vaccinations)
	assert.Equal(t, MyTime(now).DateString(), MyTime(entry.Date).DateString())
}
