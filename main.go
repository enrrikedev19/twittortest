package main

import (
	"log"

	"github.com/enrrikedev19/twittortest/bd"
	"github.com/enrrikedev19/twittortest/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexoion a la BD")
		return
	}

	handlers.Manejadores()

}
