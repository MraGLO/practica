package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MraGLO/practica/internal/app"
)

func main() {
	app := app.App{}

	app.Init("github.com/MraGLO/practica/config.yaml")

	defer app.Close()

	go app.Run()

	// handle ctr+c.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("\nServer stopped")

}
