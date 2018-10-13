package routes

import (
	"net/http"

	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/gorilla/mux"
)

func init() {

	r := mux.NewRouter()

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(static)

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(AccessLevel(aenianos.ACCESS_ADMIN))

	unlogged := r.PathPrefix("/auth").Subrouter()
	unlogged.Use(Unlogged)

	logged := r.PathPrefix("/").Subrouter()
	logged.Use(AccessLevel(aenianos.ACCESS_NORMAL))

	r.HandleFunc("/home", Home).Methods("GET").Name("home-get")

	logged.HandleFunc("/user", Home).Methods("GET").Name("user-get")
	logged.HandleFunc("/logout", Logout).Name("logout")

	unlogged.HandleFunc("/login", LoginGet).Methods("GET").Name("login-get")
	unlogged.HandleFunc("/login", LoginPost).Methods("POST").Name("login-post")

	unlogged.HandleFunc("/register", RegisterGet).Methods("GET").Name("register-get")
	unlogged.HandleFunc("/register", RegisterPost).Methods("POST").Name("register-post")

	admin.HandleFunc("/menu", AdminMenu).Methods("GET").Name("adminmenu-get")

	admin.HandleFunc("/addgenero", AddGeneroGet).Methods("GET").Name("addgenero-get")
	admin.HandleFunc("/addgenero", AddGeneroPost).Methods("POST").Name("addgenero-post")

	data.SetRouter(r)
}
