package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type UserProfile string

var db = make(map[string]string)

func pong(c *gin.Context) {
	c.String(200, "%s", "pong")
}

func main() {
	const (
		Guest UserProfile = "Guest"
		Admin UserProfile = "Admin"
	)

	/*
	returns a Gin engine instance with the Logger and Recovery middleware already attached. 
	The Logger middleware logs every request and its latency, while the Recovery middleware 
	recovers from any panics and writes a 500 error if there was one.
	*/
	router := gin.Default()

	// connect to database
	connStr := "user=root dbname=test123 sslmode=verify-full"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Database: Connection Successful")
	}

	// system health
	router.GET("/ping", pong)

	router.Run("0.0.0.0:8080")
}
