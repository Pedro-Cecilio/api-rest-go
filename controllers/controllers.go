package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Pedro-Cecilio/api-rest-go/database"
	"github.com/Pedro-Cecilio/api-rest-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func OnePersonality(w http.ResponseWriter, r *http.Request) {
	database.ConnectDB()
	var personality models.Personality
	vars := mux.Vars(r)
	id := vars["id"]
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request){
	database.ConnectDB()
	var newPersonality models.Personality
	var todasPersonalidades []models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	result := database.DB.Create(&newPersonality)
	result.Find(&todasPersonalidades)
	fmt.Println(todasPersonalidades)
	json.NewEncoder(w).Encode(newPersonality)
}
