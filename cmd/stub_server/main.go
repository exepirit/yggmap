package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	for _, endpoint := range Endpoints {
		server.Handle(endpoint.Method, endpoint.Route, endpoint.Handler)
	}

	if err := server.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
