package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	r := gin.Default()

	r.LoadHTMLFiles("templates/awesomeProject.html")
	r.Static("/static", "./static/")

	r.GET("/ping", func(c *gin.Context) {
		/* output JSON
		c.JSON(200, gin.H{
			"message": "pong",
		})
		*/
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

	log.Fatal(r.Run()) // listen and server on 0.0.0.0:8080
}
