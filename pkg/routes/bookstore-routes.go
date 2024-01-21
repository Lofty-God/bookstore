package routes

import (
	"github.com/Lofty-God/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{booId}", controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
}
