package closer

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func CloseOnSignal(signals ...os.Signal) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, signals...)

	log.Printf("got %s signal", <-done)

	return Close()
}

func CloseOnSignalWContext(ctx context.Context, signals ...os.Signal) error {
	closeCtx, _ := signal.NotifyContext(ctx, signals...)
	<-closeCtx.Done()

	log.Printf("got %s signal", closeCtx)

	return Close()
}

func CloseOnSignalContext(signals ...os.Signal) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return CloseOnSignalWContext(ctx, signals...)
	}
}
