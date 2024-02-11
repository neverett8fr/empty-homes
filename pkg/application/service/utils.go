package service

import (
	"database/sql"
	"empty-homes/pkg/infra/db"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	DBConn *db.DBConn
)

func NewServiceRoutes(r *mux.Router, conn *sql.DB) {

	DBConn = db.NewDBConnFromExisting(conn)

	newHomesInformation(r)

}

func writeReponse(w http.ResponseWriter, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
