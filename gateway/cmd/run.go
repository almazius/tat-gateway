package main

import (
	"context"
	"fmt"
	"gateway/gateway/config"
	"gateway/gateway/src/handler"
)

func run(ctx context.Context) error {
	var err error

	srv := handler.NewServer()
	err = srv.ListenAndServe(config.C().Server.Host, config.C().Server.Port)
	if err != nil {
		return fmt.Errorf("failed start server: %w", err)
	}

	return nil
}
