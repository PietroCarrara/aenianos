//
// main.go
// Copyright (C) 2018 pietro <pietro@the-arch>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/middleware"
	"github.com/PietroCarrara/aenianos/internal/routes"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	r.StrictSlash(true)

	r.HandleFunc("/home", routes.Home).Methods("GET")

	r.HandleFunc("/user", middleware.AccessLevel(aenianos.ACCESS_NORMAL, routes.Home)).Methods("GET")

	r.HandleFunc("/login", middleware.Unlogged(routes.LoginGet)).Methods("GET")
	r.HandleFunc("/login", middleware.Unlogged(routes.LoginPost)).Methods("POST")

	r.HandleFunc("/register", middleware.Unlogged(routes.RegisterGet)).Methods("GET")
	r.HandleFunc("/register", middleware.Unlogged(routes.RegisterPost)).Methods("POST")

	r.HandleFunc("/admin/addgenero", middleware.AccessLevel(aenianos.ACCESS_ADMIN, routes.AddGeneroGet)).Methods("GET")
	r.HandleFunc("/admin/addgenero", middleware.AccessLevel(aenianos.ACCESS_ADMIN, routes.AddGeneroPost)).Methods("POST")

	r.HandleFunc("/logout", middleware.AccessLevel(aenianos.ACCESS_NORMAL, routes.Logout))

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(static)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, r)
}
