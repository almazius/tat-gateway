package main

import (
	"gateway/src/config"
	"log/slog"
	"os"
)

func init() {
	// Set json format for logs
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// load and set config
	v, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config",
			slog.Any("error", err))
		panic(err)
	}

	cfg, err := config.ParseConfig(v)
	if err != nil {
		slog.Error("failed to parse config",
			slog.Any("error", err))
		panic(err)
	}

	_ = cfg
}
