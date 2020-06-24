package bd

import (
	"context"
	"time"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultarRelacion consulta una relacion*/
func ConsultarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("red-social")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return false, err
	}
	return true, nil
}
