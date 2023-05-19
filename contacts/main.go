package main

import (
	database "contacts/databases"
	"contacts/handlers"
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	dsn  string
	port string
)

func main() {

	flag.StringVar(&dsn, "dsn", "host=localhost user=postgres password=postgres dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "--dsn=host=localhost user=postgres password=postgres dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	flag.StringVar(&port, "port", "8080", "--port=8080 or -port=8080")
	flag.Parse()

	router := gin.Default()

	//dsn := "host=localhost user=postgres password=postgres dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := database.GetConnection(dsn)
	if err != nil {
		log.Fatal(err)
	}

	contactdb := database.Contact{DB: db}

	//http.ListenAndServe(":"+port, nil)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	chandler := new(handlers.Contact)
	chandler.IContact = &contactdb

	router.POST("/v1/private/contact", chandler.Create)

	router.POST("/v1/private/contact/ms", chandler.CreateUseMessageBroker)

	router.GET("v1/private/contact/:id", chandler.GetByID())

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
