package routes

import (
	"go_rest_api/controllers"
	"go_rest_api/database"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// This function is used to initialize the router and routes for it.
func InitializeRoutes(app *database.App) *database.App {
	// Initialize the router
	app.Router = mux.NewRouter()

	// Set the HTTP method for each call
	app.Router.HandleFunc("/products", handleRequest(app, controllers.GetProducts)).Methods("GET")
	app.Router.HandleFunc("/product/{id:[0-9]+}", handleRequest(app, controllers.GetProduct)).Methods("GET")
	app.Router.HandleFunc("/product", handleRequest(app, controllers.CreateProduct)).Methods("POST")
	app.Router.HandleFunc("/product/{id:[0-9]+}", handleRequest(app, controllers.UpdateProduct)).Methods("PUT")
	app.Router.HandleFunc("/product/{id:[0-9]+}", handleRequest(app, controllers.DeleteProduct)).Methods("DELETE")

	return app
}

// This function is to pass db arg to all endpoints
type RequestHandlerFunction func(db *pgx.Conn, w http.ResponseWriter, r *http.Request)

// This function acts a middleware between the request and DB.
func handleRequest(app *database.App, handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}
