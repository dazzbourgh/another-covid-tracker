package service

import (
	"another-covid-tracker.com/chart/handler"
	"another-covid-tracker.com/chart/types"
	"another-covid-tracker.com/chart/util"
	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gziphandler"
	"net/http"
)

type ChartService struct {
	URL string
}

func (s *ChartService) Prefix() string {
	return "/svc/chart"
}

func (s *ChartService) Middleware(handler http.Handler) http.Handler {
	return gziphandler.GzipHandler(handler)
}

func (s *ChartService) Endpoints() map[string]map[string]http.HandlerFunc {
	return map[string]map[string]http.HandlerFunc{
		"/{isoCode}": {
			"GET": handler.ChartHandler{Fetch: util.FetchEntries(s.URL)}.BuildChart,
		},
	}
}

type GetEntries func(isoCode string) types.Entries

type Config struct {
	Server         *server.Config
	DataServiceUrl string `envconfig:"DATA_SERVICE_URL"`
	Port           string `envconfig:"HTTP_PORT"`
}

func NewChartService(cfg *Config) *ChartService {
	return &ChartService{URL: cfg.DataServiceUrl}
}
