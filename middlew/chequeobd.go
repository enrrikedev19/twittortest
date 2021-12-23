package middlew

import (
	"net/http"

	"github.com/enrrikedev19/twittortest/bd"
)

/*ChequeoBD es el middlew que permite conocer el estado de la base datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdidad con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
