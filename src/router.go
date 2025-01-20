package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Crea un router Gin
	r := gin.Default()
	r.LoadHTMLFiles("C:/Users/zderk/Downloads/GMS/Lenguages/Go/Go-explorer/Github-up/Portfolio/src/pages/index.html", "C:/Users/zderk/Downloads/GMS/Lenguages/Go/Go-explorer/Github-up/Portfolio/src/pages/golang.html")
	r.Static("/bulma", "C:/Users/zderk/Downloads/GMS/Lenguages/Go/Go-explorer/Github-up/Portfolio/bulma")
	r.StaticFile("/styles.css", "C:/Users/zderk/Downloads/GMS/Lenguages/Go/Go-explorer/Github-up/Portfolio/src/pages/styles.css")
	r.Static("/resources", "C:/Users/zderk/Downloads/GMS/Lenguages/Go/Go-explorer/Github-up/Portfolio/src/pages/resources")
	// Ruta de home expone el index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/golang", func(c *gin.Context) {
		c.HTML(200, "golang.html", nil)
	})

	// Inicia el servidor en el puerto 8080
	r.Run(":3000")
}
