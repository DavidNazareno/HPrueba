package main

import (
	"github.com/DavidNazareno/h_prueba/controller"
	"github.com/DavidNazareno/h_prueba/handler"
)

func main() {

	controller.LoadDatabase() // Cargar base de datos

	handler.HandlerController() //Crear servidor junto a rutas
}
