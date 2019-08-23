package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/bassaer/go-user-api/app/pkg"
)

func main() {
	mux := http.NewServeMux()
	repo, err := app.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	mux.Handle("/", app.NewHandler(repo))

	svr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
