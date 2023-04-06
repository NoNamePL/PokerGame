package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Welcome struct {
	Name string
	Time string
}

//GET

func handlerLogin(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"status": http.StatusOK,
		},
	)
}

func handlerRegister(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"register.html",
		gin.H{
			"status": http.StatusOK,
		},
	)
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

func handlerMain(c *gin.Context, db *sql.DB) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"status": http.StatusOK,
		},
	)
}

// POST

func postLogin(c *gin.Context) {

}

func main() {

	//connecting := "postgres://postgres:postgrespw@localhost:32768/university?sslmode=disable"
	// connecting to DB
	//db, err := sql.Open("postgres", connecting)
	//if err != nil {
	//	log.Fatal(err)
	//}

	db := sql.DB{}

	r := gin.Default()
	//r.LoadHTMLFiles("templates/awesomeProject.html", "templates/index.html")
	r.LoadHTMLGlob("templates/*.html")

	r.Static("/static", "./static/")

	r.GET("/login", handlerLogin)
	r.GET("/register", handlerRegister)
	r.GET("/ping", handlerPing)
	r.POST("/login", postLogin)
	r.GET("/", func(ctx *gin.Context) {
		handlerMain(ctx, &db)
	})
	log.Fatal(r.Run()) // listen and server on 0.0.0.0:8080
}
