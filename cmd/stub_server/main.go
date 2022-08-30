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

	if err := server.Run("localhost:8000"); err != nil {
		log.Fatalln(err)
	}
}
