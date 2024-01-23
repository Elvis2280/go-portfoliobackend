package services

import (
	"Portfoliobackend/models"
	"Portfoliobackend/utils"
	"encoding/json"
	"log"
	"net/http"
)

func GetTags(response http.ResponseWriter, request *http.Request) {
	var tags []models.Tag    // Create a new tag struct
	db, err := utils.GetDb() // Get database connection

	if err != nil { // If there is an error, log it
		log.Fatalf("Could not connect to DB: %v", err)
	}

	listTags := db.Model(tags).Find(&tags)

	if listTags.Error != nil { // If there is an error, log it
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Error fetching tags"))
		return
	}

	listJson, errJson := json.Marshal(listTags) // Convert to JSON
	if errJson != nil {                         // If there is an error, log it
		log.Fatalf("Could not convert to JSON: %v", errJson)
	}

	response.WriteHeader(http.StatusOK)
	response.Write(listJson)
}
