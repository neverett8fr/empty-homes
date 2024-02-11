package service

import (
	"database/sql"
	"empty-homes/pkg/infra/db"
	"encoding/json"
	"fmt"
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

func homesObjectToHTML(properties []db.Home) string {
	out := ""
	for _, val := range properties {
		out += fmt.Sprintf("<div><b>%s</b><p>%s</p><p>%s</p><p>%s</p></div>", val.Name,
			fmt.Sprintf("Postcode: %s, Street: %s", val.Postcode, val.Street),
			fmt.Sprintf("Type: %s, Bedrooms: %v, Safe: %v", val.Type, val.Bedrooms, val.Safe),
			fmt.Sprintf("Date Added: %v, Last Checked: %v", val.DateAdded, val.DateLastChecked))
	}

	return out
}
