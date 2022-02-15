package webserver_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/ONSDigital/blaise-go-fizzbuzz/webserver"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FizzBuzz Controller", func() {
	var (
		httpRouter         *gin.Engine
		fizzBuzzController = &webserver.FizzBuzzController{}
	)

	BeforeEach(func() {
		httpRouter = gin.Default()
		fizzBuzzController.AddRoutes(httpRouter)
	})

	Describe("GET /fizzbuzz/:number", func() {
		var (
			httpRecorder  *httptest.ResponseRecorder
			fizzBuzzQuery string
		)

		JustBeforeEach(func() {
			httpRecorder = httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/fizzbuzz/%s", fizzBuzzQuery), nil)
			httpRouter.ServeHTTP(httpRecorder, req)
		})

		Context("when a number is provided", func() {
			BeforeEach(func() {
				fizzBuzzQuery = "3"
			})

			It("returns a Fizz http response", func() {
				Expect(httpRecorder.Code).To(Equal(http.StatusOK))
				Expect(httpRecorder.Body.String()).To(Equal(`"Fizz"`))
			})
		})

		Context("when a number is not provided", func() {
			BeforeEach(func() {
				fizzBuzzQuery = "fwibble"
			})

			It("returns an error http response", func() {
				Expect(httpRecorder.Code).To(Equal(http.StatusBadRequest))
				Expect(httpRecorder.Body.String()).To(Equal(`{"error":"fwibble is not a number"}`))
			})
		})
	})
})
