package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	r := gin.Default()

	r.LoadHTMLFiles("templates/awesomeProject.html", "templates/index.html")

	r.Static("/static", "./static/")

	r.GET("/ping", func(c *gin.Context) {

		if name := c.Request.FormValue("name"); name != "" {
			welcome.Name = name
		}
		c.HTML(
			http.StatusOK,
			"awesomeProject.html",
			gin.H{
				"status": http.StatusOK,
				"Name":   welcome.Name,
				"Time":   welcome.Time,
			},
		)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"status": http.StatusOK,
			},
		)
	})
	log.Fatal(r.Run()) // listen and server on 0.0.0.0:8080
}
