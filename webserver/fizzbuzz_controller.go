package webserver

import (
	"net/http"
	"strconv"

	"github.com/ONSDigital/blaise-go-fizzbuzz/fizzbuzz"
	"github.com/gin-gonic/gin"
)

type FizzBuzzController struct{}

func (fizzBuzzController *FizzBuzzController) AddRoutes(httpRouter *gin.Engine) {
	httpRouter.GET("/fizzbuzz/:number", fizzBuzzController.FizzBuzz)
}

func (fizzBuzzController *FizzBuzzController) FizzBuzz(context *gin.Context) {
	number := context.Param("number")

	numberAsInt, err := strconv.Atoi(number)

	result := fizzbuzz.FizzBuzz(numberAsInt)
	if err == nil {
		context.JSON(http.StatusOK, result)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "fwibble is not a number"})
	}
}
