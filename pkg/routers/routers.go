package routers

import (
	"api-books/pkg/handlers"
	"api-books/pkg/middleware"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// middleware declaration
	router.Use(middleware.JsonHeader)

	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books", handlers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", handlers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", handlers.DeleteBook).Methods("DELETE")

	return router
}