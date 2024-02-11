package service

import (
	application "empty-homes/pkg/application/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func newHomesInformation(r *mux.Router) {

	subRouter := r.PathPrefix("/homes").Subrouter()
	subRouter.HandleFunc("/test", testHandler)

}

func testHandler(w http.ResponseWriter, r *http.Request) {

	body := application.NewResponse("congrats, this page is running!")

	writeReponse(w, r, body)
}
