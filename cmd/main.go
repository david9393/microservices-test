package main

import (
	"log"

	"github.com/david9393/microservices-test/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}

}
