package main

import (
	"PW/library"
	"log"
	"net/http"

	cfg "PW/config"
	ctr "PW/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	config = cfg.Config{}
	db     = library.UserCtr{}
)

func init() {
	config.Init()
	db.Server = config.Server
	db.Database = config.Database
	db.Connect()
}

func main() {
	r := mux.NewRouter()

	guru(r)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "PUT", "POST", "HEAD", "OPTIONS"})

	if err := http.ListenAndServe(":7000", handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatal(err)
	}
}

func guru(r *mux.Router) {
	r.HandleFunc("/guru", ctr.SemuaGuru).Methods("GET")
	r.HandleFunc("/guru/{id}", ctr.CariGuru).Methods("GET")
	r.HandleFunc("/guru", ctr.DaftarGuru).Methods("POST")
	r.HandleFunc("/guru/{id}", ctr.UpdateGuru).Methods("PUT")
	r.HandleFunc("/guru/{id}", ctr.DeleteGuru).Methods("DELETE")
}
