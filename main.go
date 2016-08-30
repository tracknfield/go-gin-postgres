package main

import (
	"log"
	"os"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	
	db, err := sql.Open("postgres", "user=arzynik dbname=test sslmode=disable")
	db.SetMaxOpenConns(20)
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/", func(c *gin.Context) {
		data, jsonErr := json.Marshal(c.Request.URL.Query())
		
		if jsonErr != nil {
			c.String(400, "Error parsing data")
			return
		}

		db.Exec("insert into track (data) values($1)", data)
		c.String(200, "")
	})

	router.Run(":" + port)
}
