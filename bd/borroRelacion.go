package bd

import (
	"context"
	"time"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/models"
)

/*BorroRelacion elimina una relacion*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("red-social")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
