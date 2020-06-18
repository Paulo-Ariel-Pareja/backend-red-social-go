package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
)

/*ModificarPerfil modifica el perfil*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos", http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificarRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al actualizar el usuario", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se actualizo los datos", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
