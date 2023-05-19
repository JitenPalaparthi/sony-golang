package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	port string
)

func main() {

	flag.StringVar(&port, "port", "8081", "--port=8081 or -port=8081")
	flag.Parse()
	router := gin.Default()

	router.GET("/v1/private/service/email/:id", func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.String(http.StatusBadRequest, "invalid id parameter")
			ctx.Abort()
		}
		ctx.JSON(200, map[string]string{"email": "Sending email to id:" + id})
	})

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
