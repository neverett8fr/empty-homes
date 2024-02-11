package service

import (
	application "empty-homes/pkg/application/entities"
	"empty-homes/pkg/infra/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newHomesInformation(r *mux.Router) {

	subRouter := r.PathPrefix("/homes").Subrouter()
	subRouter.HandleFunc("/test", testHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("", newPropertyHandler).Methods(http.MethodPost)

}

func testHandler(w http.ResponseWriter, r *http.Request) {

	body := application.NewResponse("congrats, this page is running!")

	writeReponse(w, body)
}

func newPropertyHandler(w http.ResponseWriter, r *http.Request) {

	bodyIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)

		log.Printf("error reading body, err %v", err)
		return
	}

	newProperty := db.Home{}
	err = json.Unmarshal(bodyIn, &newProperty)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)

		log.Printf("error unmarshalling body, err %v", err)
		return
	}

	err = DBConn.AddProperty(newProperty)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	body := application.NewResponse(fmt.Sprintf("property created with name %v", newProperty.Name), err)

	writeReponse(w, body)
}
