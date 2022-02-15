package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FizzBuzzController struct{}

func (fizzBuzzController *FizzBuzzController) AddRoutes(httpRouter *gin.Engine) {
	httpRouter.GET("/fizzbuzz/:number", fizzBuzzController.FizzBuzz)
}

func (fizzBuzzController *FizzBuzzController) FizzBuzz(context *gin.Context) {
	number := context.Param("number")
	context.JSON(http.StatusOK, number)
}
