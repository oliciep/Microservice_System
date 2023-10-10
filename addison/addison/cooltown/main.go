package main

// Importing packages and modules from resources file
import (
	"log"
	"net/http"
	"cooltown/resources"
)

// Main function starts listening on port 3002
func main() {
	log.Fatal(http.ListenAndServe(":3002", resources.Router()))
}
