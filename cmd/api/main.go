package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/picolloo/auth-playground/cmd/api/handler"
	"github.com/picolloo/auth-playground/infra/inmem"
	"github.com/picolloo/auth-playground/usecase/user"
)

func main() {
	r := mux.NewRouter()

	user_repo := inmem.NewUserRepository()
	user_service := user.NewService(user_repo)
	r.HandleFunc("/users", handler.GetUsers(user_service)).Methods("GET")
	r.HandleFunc("/users", handler.CreateUser(user_service)).Methods("POST")

	http.Handle("/", r)

	http.ListenAndServe(":3000", http.DefaultServeMux)
}