package main

import (
	"log"
	"net/http"

	ctr "PW/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	guru(r)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "DELETE", "PUT"})

	if err := http.ListenAndServe(":7000", handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatal(err)
	}
}

func guru(r *mux.Router) {
	r.HandleFunc("/guru", ctr.SemuaGuru).Methods("GET")
	r.HandleFunc("/guru/{id}", ctr.CariGuru).Methods("GET")
	r.HandleFunc("/guru", ctr.DaftarGuru).Methods("POST")
	r.HandleFunc("/guru/{id}", ctr.UpdateGuru).Methods("PUT")
	r.HandleFunc("/guru", ctr.DeleteGuru).Methods("DELETE")
}
