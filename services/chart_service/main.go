package main

import (
	"another-covid-tracker.com/chart/service"
	"github.com/NYTimes/gizmo/server"
	"github.com/kelseyhightower/envconfig"
	"strconv"
)

func main() {
	var cfg service.Config
	envconfig.Process("", &cfg)
	cfg.Server = &server.Config{}
	envconfig.Process("", &cfg.Server)
	cfg.Server.HTTPPort, _ = strconv.Atoi(cfg.Port)

	server.Init("another-covid-tracker-json-proxy", cfg.Server)

	err := server.Register(service.NewChartService(&cfg))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

	err = server.Run()
	if err != nil {
		server.Log.Fatal("server encountered a fatal error: ", err)
	}
}
