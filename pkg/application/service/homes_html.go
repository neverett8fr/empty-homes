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
	standardHTML = "<html><head><style></style></head><body>|fn_nav|</body></html>"
)

func newHomesHTML(r *mux.Router) {

	subRouter := r.PathPrefix("/homes/html").Subrouter()
	subRouter.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/all", allPropertiesHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/add_form", addPropertyFormHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/add", addPropertyHandler).Methods(http.MethodGet)

	_ = HTMLProv.AddComponent("/", confHTML.Style, standardHTML, map[string]string{"|fn_nav|": html.NavBody()})

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte(HTMLProv.WebComponents["/"].Compile()))
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}

func addPropertyFormHandler(w http.ResponseWriter, r *http.Request) {

	style := "<style>.container{border:2px solid black;margin:10px}</style>"
	elements := `
	Add a property:
	<div class="container">
	<form action="/homes/html/add">
	<label for="name">Name:</label><br>
	<input type="text" id="name" name="name"><br>

	<label for="postcode">Postcode:</label><br>
	<input type="text" id="postcode" name="postcode"><br>

	<label for="street">Street:</label><br>
	<input type="text" id="street" name="street"><br>

	<label for="type">Type:</label><br>
	<input type="text" id="type" name="type"><br>

	<label for="bedrooms">Bedrooms:</label><br>
	<input type="number" id="bedrooms" name="bedrooms"><br>

	<label for="inhabited">Inhabited:</label><br>
	<input type="checkbox" id="inhabited" name="inhabited"><br>

	<label for="safe">Safe:</label><br>
	<input type="checkbox" id="safe" name="safe"><br>

	<label for="owner-name">Owner Name:</label><br>
	<input type="text" id="owner-name" name="owner-name"><br>

	<label for="owner_contact">Owner Contact Info:</label><br>
	<input type="text" id="owner_contact" name="owner_contact"><br>

	<label for="last_checked">Last Checked:</label><br>
	<input type="datetime-local" id="last_checked" name="last_checked"><br>

	<input type="submit" value="Submit">
	</form>
	</div>
	`

	html := fmt.Sprintf("<html><head>%s</head><body>%s</body></html>", style, elements)

	_, err := w.Write([]byte(html))
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
