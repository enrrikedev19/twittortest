package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/enrrikedev19/twittortest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil bsuco perfil en la bd*/
func BuscoPerfil(ID string) (models.Usuario, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	bd := MongoCN.Database("twittortest")
	col := bd.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil
}
