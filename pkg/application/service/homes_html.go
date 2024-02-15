package service

import (
	"empty-homes/pkg/infra/db"
	"empty-homes/pkg/infra/html"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	standardHTML = "<html><head><style></style></head><body>|fn_nav|%s</body></html>"
)

func newHomesHTML(r *mux.Router) {

	subRouter := r.PathPrefix("/homes/html").Subrouter()
	subRouter.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/all", allPropertiesHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/add_form", addPropertyFormHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/add", addPropertyHandler).Methods(http.MethodGet)

	_ = HTMLProv.AddComponent("/", confHTML.Style, fmt.Sprintf(standardHTML, ""),
		map[string]string{"|fn_nav|": html.NavBody()})
	_ = HTMLProv.AddComponent("/add_form", confHTML.Style, fmt.Sprintf(standardHTML, "|fn_add_form|"),
		map[string]string{"|fn_nav|": html.NavBody(), "|fn_add_form|": html.PropertyFormBody()})

	prop, _ := DBConn.ViewAll() // this is only compiled at start, and not on change
	_ = HTMLProv.AddComponent("/all", confHTML.Style, fmt.Sprintf(standardHTML, "|fn_all|"),
		map[string]string{"|fn_nav|": html.NavBody(), "|fn_all|": html.AllPropertiesBody(prop)})

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte(HTMLProv.WebComponents["/"].Compile()))
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}

func addPropertyFormHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte(HTMLProv.WebComponents["/add_form"].Compile()))
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}

func addPropertyHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	postcode := r.URL.Query().Get("postcode")
	street := r.URL.Query().Get("street")
	ty := r.URL.Query().Get("type")
	bedrooms := r.URL.Query().Get("bedrooms")
	inhabited := r.URL.Query().Get("inhabited")
	safe := r.URL.Query().Get("safe")
	ownerName := r.URL.Query().Get("owner_name")
	ownerContact := r.URL.Query().Get("owner_contact")
	lastChecked := r.URL.Query().Get("last_checked")

	bR, err := strconv.Atoi(bedrooms)
	if err != nil {
		log.Printf("error converting type, err %v", err)
		bR = 0
	}
	in := false
	if inhabited != "" {
		in = true
	}
	saf := false
	if safe != "" {
		saf = true
	}
	lC, err := time.Parse(time.RFC3339Nano, lastChecked)
	if err != nil {
		log.Printf("error converting type, err %v", err)
	}

	h := db.Home{
		Name: name, Postcode: postcode, Street: street, Type: ty, Bedrooms: bR, Inhabited: in,
		Safe: saf, OwnerName: ownerName, OwnerContact: ownerContact, DateLastChecked: lC,
	}

	err = DBConn.AddProperty(h)
	if err != nil {
		log.Printf("error adding property, err %v", err)
	}

	allPropertiesHandler(w, r)

}

func allPropertiesHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte(HTMLProv.WebComponents["/all"].Compile()))
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
