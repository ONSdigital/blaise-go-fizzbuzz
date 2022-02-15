package main

import (
	"fmt"

	"github.com/ONSDigital/blaise-go-fizzbuzz/webserver"
)

func main() {
	server := &webserver.Server{}
	httpRouter := server.SetupRouter()
	httpRouter.Run(fmt.Sprintf(":9999"))
}
