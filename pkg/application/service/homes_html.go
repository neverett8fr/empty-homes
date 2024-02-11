package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newHomesHTML(r *mux.Router) {

	subRouter := r.PathPrefix("/homes/html").Subrouter()
	subRouter.HandleFunc("/", allPropertiesHandler).Methods(http.MethodGet)

}

func allPropertiesHandler(w http.ResponseWriter, r *http.Request) {

	html := "<html><head></head><body>"
	properties, err := DBConn.ViewAll()
	if err == nil {
		html += homesObjectToHTML(properties)
	} else {
		html += "<div> there has been an error </div>"
	}

	html += "</body></html>"

	_, err = w.Write([]byte(html))
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
