package middlewares

import (
	"net/http"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
)

/*ChequeoBD verifica conexion a base de datos y retorna al siguiente middleware*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
