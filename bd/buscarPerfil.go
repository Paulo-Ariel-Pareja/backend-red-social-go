package bd

import (
	"context"
	"log"
	"time"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscarPerfil busca un perfil en la bd*/
func BuscarPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("red-social")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	perfil.Password = ""
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	if err != nil {
		log.Fatal(err.Error())
		return perfil, err
	}
	return perfil, nil
}
