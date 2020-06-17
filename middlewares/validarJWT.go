package middlewares

import (
	"net/http"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/routes"
)

/*ValidarJWT verifica jwt*/
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token invalido", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
