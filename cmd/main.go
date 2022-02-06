package main

import (
	"log"

	"gomod.pekora.dev/tabasco/internal/interface/httpapi"
)

func main() {
	log.Println("Tabasco in your mouse!!")

	httpapi.LunchHTTPServer()
}
