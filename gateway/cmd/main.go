package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	defer cancel()

	err := run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to run program",
			slog.Any("error", err))
		return
	}

	<-ctx.Done()
	gracefulStop()
}

func gracefulStop() {

}
