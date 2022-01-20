package infrastructure

import "github.com/gin-gonic/gin"

type Server struct {
	Gin *gin.Engine
}

func NewServer(logger Logger) Server {
	engine := gin.New()
	engine.Use(gin.Recovery(), logger.GetGinLogger())
	return Server{Gin: engine}
}
