package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/jwt"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
)

/*Login realiza un login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña incorrectos "+err.Error(), 403)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es obligatorio", 400)
		return
	}
	user, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña incorrectos", 403)
		return
	}
	jwtKey, err := jwt.GeneroJWT(user)
	if err != nil {
		http.Error(w, "Error generando token"+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
