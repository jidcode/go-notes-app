package main

import (
	"log"
	"net/http"

	config "github.com/code/go-project/db"
	"github.com/code/go-project/handlers"
	"github.com/code/go-project/repository"
	"github.com/code/go-project/routes"
)

// Entry point of the app
func main() {
	// Connect to the database
	db := config.ConnectDB()

	// Set up repository and handler
	noteRepo := repository.NewNoteRepository(db)
	noteHandler := handlers.NewNoteHandler(noteRepo)

	// Setup routes
	router := routes.SetupRoutes(noteHandler)

	// Start the server on port 5000
	log.Println("Starting server on :5000...")
	http.ListenAndServe(":5000", router)
}
