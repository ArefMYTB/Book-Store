package routes

import (
	"Bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}
