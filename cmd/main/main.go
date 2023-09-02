package main

import (
	"log"
	"os"
	"os/signal"
	api "personal/go-htmx/cmd/web/routes"
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

	api.Run(killSignal) // Blocks anything after this
}
