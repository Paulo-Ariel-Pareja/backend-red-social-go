package main

import (
	"log"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("No se pudo conectar a la db")
		return
	}
	handlers.Manejadores()
}
