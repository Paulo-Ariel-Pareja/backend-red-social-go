package bd

import (
	"log"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza verificacion de login en bd*/
func IntentoLogin(email, password string) (models.Usuario, bool) {
	user, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		log.Fatal(err.Error())
		return user, false
	}
	return user, true
}
