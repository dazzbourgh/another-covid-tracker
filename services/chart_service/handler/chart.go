package handler

import (
	"another-covid-tracker.com/chart/util"
	"fmt"
	"github.com/NYTimes/gizmo/server"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	chartTypes "github.com/go-echarts/go-echarts/v2/types"
	"net/http"
	"strings"
)

type ChartHandler struct {
	Fetch util.FetchEntriesFunc
}

func (c ChartHandler) BuildChart(w http.ResponseWriter, r *http.Request) {
	isoCode := server.Vars(r)["isoCode"]
	from := util.TimeString(r.URL.Query().Get("from")).Date()
	to := util.TimeString(r.URL.Query().Get("to")).Date()
	lineChart := charts.NewLine()
	lineChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: chartTypes.ThemeInfographic}),
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("Vaccinations/Cases for %s", strings.ToUpper(isoCode)),
		}))
	entries := c.Fetch(isoCode, from, to)
	cases := make([]float32, len(entries))
	vaccinations := make([]float32, len(entries))
	for i := range cases {
		cases[i] = entries[i].Cases
		vaccinations[i] = entries[i].Vaccinations
	}
	lineChart.
		SetXAxis(vaccinations).
		AddSeries("Cases", generateLineItems(cases)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{}))
	err := lineChart.Render(w)
	if err != nil {
		server.Log.Debug(err)
	}
}

func generateLineItems(cases []float32) []opts.LineData {
	items := make([]opts.LineData, len(cases))
	for i := range cases {
		items[i] = opts.LineData{Value: cases[i]}
	}
	return items
}
