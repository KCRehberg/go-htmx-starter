package routes

import (
	"go-htmx/cmd/web/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(os_signal chan os.Signal) {
	router := gin.Default()

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	router.LoadHTMLGlob(pwd + "/internal/src/templates/**/*.html")
	router.Static("/static", pwd+"/internal/src/static")
	router.Static("/assets", pwd+"/internal/src/public")

	controllers.Pages(router)
	controllers.Api(router)

	// Start http server & block
	go router.Run()
	<-os_signal
	log.Println("Shutdown server...")
}
