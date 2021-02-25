package util

import (
	"another-covid-tracker.com/chart/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type FetchEntriesFunc func(string, time.Time, time.Time) types.Entries

func FetchEntries(url string) FetchEntriesFunc {
	return func(isoCode string, from time.Time, to time.Time) types.Entries {
		req, _ := http.NewRequest("GET", url, nil)
		req.URL.Path = fmt.Sprintf("/%s", isoCode)
		q := req.URL.Query()
		q.Add("from", MyTime(from).DateString())
		q.Add("to", MyTime(to).DateString())
		req.URL.RawQuery = q.Encode()
		resp, err := http.Get(req.URL.String())
		if err != nil {
			return make(types.Entries, 0)
		} else {
			defer resp.Body.Close()
			var result types.Entries
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(bodyBytes, &result)
			return result
		}
	}
}
