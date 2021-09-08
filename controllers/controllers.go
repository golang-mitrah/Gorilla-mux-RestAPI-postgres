package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	"go_rest_api/models"
)

func GetProducts(db *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	// Converting from string into int
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	// If no count value given, set it as 10 by default
	if count <= 0 {
		count = 10
	}
	fmt.Printf("%d", count)
	fmt.Printf("%d", start)
	// If no start value given, set it as 0 by default
	if start < 0 {
		start = 0
	}

	products, err := models.GetProducts(db, start, count)
	if err != nil {
		// If error, return as internal server error
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func GetProduct(db *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	// Retrieve variables from the request payload
	vars := mux.Vars(r)

	// Converting from string into int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If error return with message
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := models.Product{ID: id}
	if err := p.GetProduct(db); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func CreateProduct(db *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	var p models.Product
	// Creating a log file to check
	file, err := os.OpenFile("gorilla-mux-restapi-postgres.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// Read data from the request payload ( input stream )
	decoder := json.NewDecoder(r.Body)

	// Set the output file for logging
	log.SetOutput(file)

	// Decoding the data and put into the product reference variable
	if err := decoder.Decode(&p); err != nil {
		// Logging if there is an err
		log.Fatal(err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	// Delaying the execution and closing it when the need is over
	defer r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := p.CreateProduct(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

// This function is used to update a product by giving the unique identifier (ID)
func UpdateProduct(db *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	// Retrieve variables from the request payload
	vars := mux.Vars(r)

	// Converting from string into int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p models.Product

	// Read data from the request payload ( input stream )
	decoder := json.NewDecoder(r.Body)

	// Decoding the data and put into the product reference variable
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Delaying the execution and closing it when the need is over
	defer r.Body.Close()
	p.ID = id

	if err := p.UpdateProduct(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

// This function is used to delete a product by giving the unique identifier (ID)
func DeleteProduct(db *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := models.Product{ID: id}
	if err := p.DeleteProduct(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// This function is used to return the error with message in JSON format
func respondWithError(w http.ResponseWriter, code int, message string) {
	// sending it as key-value pair using map
	respondWithJSON(w, code, map[string]string{"error": message})
}

// This function is used to return the payload to the user in JSON format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Converting the payload to JSON
	response, _ := json.Marshal(payload)

	// Set the custom response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
