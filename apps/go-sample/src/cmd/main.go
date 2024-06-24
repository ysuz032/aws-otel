package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"go-sample/pkg/otel"
	"go-sample/pkg/web"
	_ "net/http/pprof"
)

func main() {
	// Stop if press Ctrl+C
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	appName := "go-sample"
	appVersion := "1.0.0"
	shutdown, err := otel.Setup(ctx, appName, appVersion)
	if err != nil {
		fmt.Println("Failed to initialize OpenTelemetry:", err)
		return
	}
	defer shutdown(ctx)

	router := web.NewRouter()

	var wg sync.WaitGroup
	wg.Add(1)

	// run pprof
	go func() {
		log.Println("Starting pprof server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof server failed: %v", err)
		}
	}()

	// run web server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
