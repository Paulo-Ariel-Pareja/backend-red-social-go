package routes

import (
	"errors"
	"strings"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/bd"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
	"github.com/dgrijalva/jwt-go"
)

/*Email email del usuario*/
var Email string

/*IDUsuario id del usuario*/
var IDUsuario string

/*ProcesarToken procesa el token jwt*/
func ProcesarToken(tk string) (*models.Claim, bool, string, error) {
	clavePrivada := []byte("EstaCl4vEesPriv@d@.")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return clavePrivada, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
