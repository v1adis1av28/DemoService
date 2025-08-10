package main

import (
	"demo/internal/app"
	"demo/internal/database"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	db := database.NewDB("postgres://postgres:postgres@db:5432/advertisements?sslmode=disable")
	app := app.NewApp(db)

	go func() {
		app.MustStart()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.Stop()
	log.Fatal("Gracefully stopped")

}
