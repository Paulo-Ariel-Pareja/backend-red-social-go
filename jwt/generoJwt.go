package jwt

import (
	"time"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
	"github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genera jwt en base a un usuario*/
func GeneroJWT(u models.Usuario) (string, error) {
	clavePrivada := []byte("EstaCl4vEesPriv@d@.")
	payload := jwt.MapClaims{
		"email":            u.Email,
		"nombre":           u.Nombre,
		"apellidos":        u.Apellidos,
		"fecha_nacimiento": u.FechaNacimiento,
		"biografia":        u.Biografia,
		"ubicacion":        u.Ubicacion,
		"sitioweb":         u.SitioWeb,
		"_id":              u.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(clavePrivada)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
