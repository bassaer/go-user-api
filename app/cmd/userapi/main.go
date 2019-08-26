package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/bassaer/go-user-api/app/pkg"
	_ "github.com/go-sql-driver/mysql"
)

const auth = "root:test@tcp(db:3306)/userdb?parseTime=true&loc=Asia%2FTokyo"

func main() {
	db, err := sql.Open("mysql", auth)
	if err != nil {
		log.Fatal(err)
	}
	repo, err := app.NewUserRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	mux := http.NewServeMux()
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
