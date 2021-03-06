package service

import (
	"another-covid-tracker.com/chart/handler"
	"another-covid-tracker.com/chart/types"
	"another-covid-tracker.com/chart/util"
	"context"
	"fmt"
	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gziphandler"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

type ChartService struct {
	URL    string
	client *http.Client
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
			"GET": handler.ChartHandler{Fetch: util.FetchEntries(s.URL, s.client)}.BuildChart,
		},
	}
}

type GetEntries func(isoCode string) types.Entries

type Config struct {
	Server         *server.Config
	DataServiceUrl string `envconfig:"DATA_SERVICE_URL"`
	Port           string `envconfig:"HTTP_PORT"`
	KeyFile        string `envconfig:"KEY_FILE"`
}

func NewChartService(cfg *Config) *ChartService {
	ctx := context.Background()
	var options []option.ClientOption
	if cfg.KeyFile != "" {
		options = make([]option.ClientOption, 1)
		options[0] = option.WithCredentialsFile(cfg.KeyFile)
	} else {
		options = make([]option.ClientOption, 0)
	}
	client, err := idtoken.NewClient(ctx, cfg.DataServiceUrl, options...)
	if err != nil {
		log.Fatal(fmt.Sprintf("error: %s", err))
	}
	return &ChartService{URL: cfg.DataServiceUrl, client: client}
}
