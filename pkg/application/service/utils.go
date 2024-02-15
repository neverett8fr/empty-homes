package service

import (
	"database/sql"
	"empty-homes/pkg/config"
	"empty-homes/pkg/infra/db"
	"empty-homes/pkg/infra/html"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	DBConn   *db.DBConn
	HTMLProv *html.HTMLProvider
	confHTML *config.HTML
)

func NewServiceRoutes(r *mux.Router, conn *sql.DB, conf *config.HTML) {

	DBConn = db.NewDBConnFromExisting(conn)
	HTMLProv, _ = html.InitialiseProvider()

	confHTML = conf

	newHomesInformation(r)
	newHomesHTML(r)

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
