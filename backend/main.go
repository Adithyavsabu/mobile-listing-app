package main

import (
	"database/sql"
	"fmt"
	"log"
	"mobile_listing_app/controllers"
	"mobile_listing_app/database"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var db *sql.DB = database.OpenDatabase()

	defer database.CloseDatabase(db)
	controllers.SetDB(db)

	// Pass the db variable to the handlers package

	router := mux.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	//producthandler routes

	router.HandleFunc("/create", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/list", controllers.GetAllProduct).Methods("GET")
	router.HandleFunc("/list/{id}", controllers.GetOneProduct).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/search/{query}", controllers.SearchProduct).Methods("GET")

	//userhandler routes

	router.HandleFunc("/UserValidation", controllers.UserValidation).Methods("POST")
	router.HandleFunc("/register", controllers.UserRegistration).Methods("POST")
	fmt.Println("Serving on server localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))

}
