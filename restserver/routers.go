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
	/*
	   Test url:
	   localhost:8000/assets/points=6.288250,-75.612795,6.210856,-75.631614,6.202257,-75.552307,6.257931,-75.536343,6.288250,-75.612795
	*/
}
