package restserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRestServer() {
	router := mux.NewRouter()

	router.HandleFunc("/assets/points={points}", GetAssets).Methods("GET")
	//	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	//	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
