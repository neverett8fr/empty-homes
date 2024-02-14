package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newHomesHTML(r *mux.Router) {

	subRouter := r.PathPrefix("/homes/html").Subrouter()
	subRouter.HandleFunc("/", allPropertiesHandler).Methods(http.MethodGet)

}

func allPropertiesHandler(w http.ResponseWriter, r *http.Request) {

	style := "<style>.container{border:2px solid black;margin:10px}</style>"

	html := fmt.Sprintf("<html><head>%s</head><body>", style)
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
