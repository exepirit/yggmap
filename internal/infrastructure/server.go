package infrastructure

import "github.com/gin-gonic/gin"

type Server struct {
	Gin *gin.Engine
}

func NewServer() Server {
	return Server{
		Gin: gin.Default(),
	}
}
