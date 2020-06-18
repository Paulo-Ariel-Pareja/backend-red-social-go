package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
)

/*InsertarTweet inserta un tweet*/
func InsertarTweet(w http.ResponseWriter, r *http.Request) {
	var t models.Tweet

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos", http.StatusBadRequest)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: t.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Error al insertar el tweet", http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se guardo el tweet", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
