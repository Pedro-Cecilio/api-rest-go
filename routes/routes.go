package routes

import (
	"log"
	"net/http"

	"github.com/Pedro-Cecilio/api-rest-go/controllers"
	"github.com/gorilla/mux"
)
func HandleRequest(){
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.OnePersonality).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreateNewPersonality).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}