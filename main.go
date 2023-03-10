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

	r := gin.Default()
	//r.LoadHTMLFiles("templates/awesomeProject.html", "templates/index.html")
	r.LoadHTMLGlob("templates/*.html")

	r.Static("/static", "./static/")

	r.GET("/ping", handlerPing)
	r.GET("/", handlerMain)
	r.GET("/login", handlerLogin)
	r.GET("/register", handlerRegister)
	log.Fatal(r.Run()) // listen and server on 0.0.0.0:8080
}

func handlerLogin(c *gin.Context) {

}

func handlerRegister(c *gin.Context) {

}

func handlerPing(c *gin.Context) {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
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
}

func handlerMain(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"status": http.StatusOK,
		},
	)
}
