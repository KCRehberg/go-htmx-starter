package routes

import (
	"log"
	"os"
	"personal/go-htmx/cmd/web/controllers"

	"github.com/gin-gonic/gin"
)

func Run(os_signal chan os.Signal) {
	router := gin.Default()

	router.LoadHTMLGlob("/mnt/c/Users/Skaro/Desktop/workspace/personal/go-htmx/internal/templates/**/*.html")
	router.Static("/static", "/mnt/c/Users/Skaro/Desktop/workspace/personal/go-htmx/static")

	controllers.Pages(router)
	controllers.Api(router)

	// Start http server & block
	go router.Run()
	<-os_signal
	log.Println("Shutdown server...")
}
