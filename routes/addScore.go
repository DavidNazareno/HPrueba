package routes

import (
	"net/http"
	"strconv"

	"github.com/DavidNazareno/h_prueba/services"
)
//AddScore ruta para actualizar la calificacion del cliente
func AddScore(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query()["token"][0]
	score, _ := strconv.Atoi(r.URL.Query()["score"][0])

	err := services.UpdateRequest(token, score)

	if err != nil {
		http.Error(w, "Error with Update Score"+err.Error(), 400)
	}
	w.WriteHeader(200)
}
