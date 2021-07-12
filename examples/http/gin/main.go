package main

import (
	"github.com/dayu-go/gkit/transport/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})

	httpSrv := http.NewServer(http.Address(":8000"))
	httpSrv.HandlePrefix("/", router)

}
