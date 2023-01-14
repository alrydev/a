package routes

import (
	"waysbookssssssss/handlers"
	"waysbookssssssss/pkg/middleware"
	"waysbookssssssss/pkg/mysql"
	"waysbookssssssss/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/user", middleware.Auth(h.GetUserByID)).Methods("GET")
	r.HandleFunc("/user-update", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.DeleteUser)).Methods("DELETE")
}
