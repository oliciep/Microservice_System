package main

// Importing packages and modules from resources file
import (
	"log"
	"net/http"
	"search/resources"
)

// Main function starts listening on port 3001
func main() {
	log.Fatal(http.ListenAndServe(":3001", resources.Router()))
}
