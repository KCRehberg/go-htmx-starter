package routes

import (
	"log"
	"os"
	"personal/go-htmx/cmd/web/controllers"

	"github.com/gin-gonic/gin"
)

func Run(os_signal chan os.Signal) {
	router := gin.Default()

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	router.LoadHTMLGlob(pwd + "/internal/templates/**/*.html")
	router.Static("/static", pwd+"/static")

	controllers.Pages(router)
	controllers.Api(router)

	// Start http server & block
	go router.Run()
	<-os_signal
	log.Println("Shutdown server...")
}
