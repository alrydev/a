package routes

import (
	"waysbookssssssss/handlers"
	"waysbookssssssss/pkg/middleware"
	"waysbookssssssss/pkg/mysql"
	"waysbookssssssss/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
	r.HandleFunc("/user-cart", middleware.Auth(h.GetCartByTransID)).Methods("GET")
}
