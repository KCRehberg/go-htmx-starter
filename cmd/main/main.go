package main

import (
	api "go-htmx/cmd/web/routes"
	"go-htmx/internal/database"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.Println("Starting server...")
	killSignal := make(chan os.Signal, 1)
	signal.Notify(
		killSignal,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGSEGV)

	database.Connect()

	api.Run(killSignal) // Blocks anything after this
}
