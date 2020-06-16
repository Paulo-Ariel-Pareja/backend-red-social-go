package bd

import (
	"context"
	"time"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario verifica existencia de usuario en bd*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("red-social")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
