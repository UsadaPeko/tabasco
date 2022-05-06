package main

import (
	"gomod.pekora.dev/tabasco/internal/eventadapter/interfaces/apiserver"

	"log"
)

func main() {
	log.Println("Tabasco in your mouse!!")

	apiserver.StartHTTPServer()
}
