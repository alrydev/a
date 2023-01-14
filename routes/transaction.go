package routes

import (
	"waysbookssssssss/handlers"
	"waysbookssssssss/pkg/middleware"
	"waysbookssssssss/pkg/mysql"
	"waysbookssssssss/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransaction)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/my-trans", middleware.Auth(h.GetOrderByID)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
