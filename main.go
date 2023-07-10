package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("starting...")

	http.HandleFunc("/", rootHandler)

	server := http.Server{Addr: ":80"}

	go func() {
		slog.Error(
			"server error",
			"errorMsg", server.ListenAndServe(),
		)
	}()

	slog.Info(
		"app started",
		slog.String("addr", server.Addr),
	)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	switch <-signalChan {
	case syscall.SIGTERM:
		slog.Info("Received SIGTERM")
		slog.Info("Sleep 5 seconds...")
		time.Sleep(5 * time.Second)
		slog.Info("5 seconds passed")
		os.Exit(0)
	case syscall.SIGINT:
		slog.Info("Received SIGINT")
		slog.Info("Sleep 5 seconds...")
		time.Sleep(5 * time.Second)
		slog.Info("5 seconds passed")
		os.Exit(0)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
