package main

import (
	"go_rest_api/database"
	"go_rest_api/routes"
	"net/http"
)

// Main will be called by default
func main() {
	// Initialize the database connection
	app := database.InitializeApp()
	r := routes.InitializeRoutes(app)

	// Run
	http.ListenAndServe(":8010", r.Router)
}
