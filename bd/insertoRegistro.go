package bd

import (
	"context"
	"time"

	"github.com/enrrikedev19/twittortest/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la funcion que registra en base de datos al usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittortest")
	col := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
