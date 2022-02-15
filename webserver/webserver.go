package webserver

import "github.com/gin-gonic/gin"

type Server struct{}

func (server *Server) SetupRouter() *gin.Engine {
	httpRouter := gin.Default()

	fizzBuzzController := &FizzBuzzController{}
	fizzBuzzController.AddRoutes(httpRouter)

	return httpRouter
}
