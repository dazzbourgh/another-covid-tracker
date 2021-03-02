package handler

import (
	"another-covid-tracker.com/chart/types"
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
	entries := c.Fetch(isoCode, from, to)
	chartEntries := util.GetChartData(entries)
	items := generateLineItems(chartEntries)

	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: chartTypes.ThemeInfographic}),
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("Vaccinations/Cases for %s", strings.ToUpper(isoCode)),
		}),
		charts.WithXAxis3DOpts(opts.XAxis3D{Name: "Time", Type: "time", Min: chartEntries[0].Date.UnixNano() / 1000000}),
		charts.WithYAxis3DOpts(opts.YAxis3D{Name: "Vaccinations", Min: chartEntries[0].Vaccinations}),
		charts.WithZAxis3DOpts(opts.ZAxis3D{Name: "Cases"}),
	)
	line3d.AddSeries("line3D", items)
	line3d.SetGlobalOptions()
	err := line3d.Render(w)
	if err != nil {
		server.Log.Debug(err)
	}
}

func generateLineItems(chartEntries types.Entries) []opts.Chart3DData {
	data := make([][3]float64, len(chartEntries))
	for i := 0; i < len(chartEntries); i++ {
		data[i] = [3]float64{
			float64(chartEntries[i].Date.UnixNano() / 1000000),
			chartEntries[i].Vaccinations,
			chartEntries[i].Cases,
		}
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opts.Chart3DData{Value: []interface{}{d[0], d[1], d[2]}})
	}
	return ret
}
