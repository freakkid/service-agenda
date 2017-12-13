package main

import (
	"os"

	"github.com/freakkid/service-agenda/service/server"
	flag "github.com/spf13/pflag"
)

// PORT is defalut to set 8080
const PORT string = "8080"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	serverInstance := server.NewServer()
	serverInstance.Run(":" + port)
}
