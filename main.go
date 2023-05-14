package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Users struct {
	ID int ``
}

type Welcome struct {
	Name string
	Time string
}

//GET

func handlerLogin(c *gin.Context, db *sql.DB) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"status": http.StatusOK,
		},
	)
}

func handlerRegister(c *gin.Context, db *sql.DB) {
	c.HTML(
		http.StatusOK,
		"register.html",
		gin.H{
			"status": http.StatusOK,
		},
	)
}

func handlerPing(c *gin.Context, db *sql.DB) {
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

func postLogin(c *gin.Context, db *sql.DB) {

}

func postRegister(c *gin.Context, db *sql.DB) {

}

/*
const (

	host     = "localhost"
	port     = 32768
	user     = "postgres"
	password = "postgrespw"
	dbname   = "PokerGame"

)
*/
func main() {

	connecting := "postgres://postgres:postgrespw@localhost:32768/PokerGame?sslmode=disable"
	// connecting to DB
	/*
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	*/
	db, err := sql.Open("postgres", connecting)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//	db := sql.DB{}

	r := gin.Default()
	//r.LoadHTMLFiles("templates/awesomeProject.html", "templates/index.html")
	r.LoadHTMLGlob("templates/*.html")

	r.Static("/static", "./static/")

	r.GET("/login", func(ctx *gin.Context) {
		handlerLogin(ctx, db)
	})
	r.GET("/register", func(ctx *gin.Context) {
		handlerRegister(ctx, db)
	})
	r.GET("/ping", func(ctx *gin.Context) {
		handlerPing(ctx, db)
	})
	r.POST("/login", func(ctx *gin.Context) {
		postLogin(ctx, db)
	})
	r.POST("/register", func(ctx *gin.Context) {
		postRegister(ctx, db)
	})
	r.GET("/", func(ctx *gin.Context) {
		handlerMain(ctx, db)
	})

	log.Fatal(r.Run()) // listen and server on 0.0.0.0:8080
}
