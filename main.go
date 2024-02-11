package main

import (
	"database/sql"
	"empty-homes/cmd"
	"empty-homes/pkg/application/service"
	"empty-homes/pkg/config"
	"log"

	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes(conn *sql.DB) *mux.Router {
	r := mux.NewRouter()
	service.NewServiceRoutes(r, conn)

	return r
}

// Initiate web server
func main() {
	conf, err := config.Initialise()
	if err != nil {
		log.Fatalf("error initialising config, err %v", err)
		return
	}
	log.Println("config initialised")

	serviceDB, err := cmd.OpenDB(&conf.DB)
	if err != nil {
		log.Fatalf("error starting db, err %v", err)
		return
	}
	defer serviceDB.Close()
	log.Println("connection to DB setup")

	err = cmd.MigrateDB(serviceDB, conf.DB.Driver)
	if err != nil {
		log.Fatalf("error running DB migrations, %v", err)
		return
	}
	log.Println("DB migrations ran")

	router := getRoutes(serviceDB)
	log.Println("API routes retrieved")

	err = cmd.StartServer(&conf.Service, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
		return
	}
	log.Println("server started")

}
