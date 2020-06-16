package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
)

/*Registro guarda en base de datos el registro de un usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es obligatorio", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser igual o mayor a 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "El usuario ya esta registrado", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error al guardar el registro: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se guardado el registro", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
