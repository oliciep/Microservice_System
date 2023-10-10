package main

// Importing packages and modules from repository and resources files
import (
	"log"
	"net/http"
	"tracks/repository"
	"tracks/resources"
)

// Main function creates repository for database and starts listening on port 3000
func main() {
	repository.Init()
	repository.Create()
	log.Fatal(http.ListenAndServe(":3000", resources.Router()))
}
