package routes

import (
	"net/http"
	"net/url"

	"github.com/PietroCarrara/aenianos"
	"github.com/gorilla/mux"
)

func Redirect(w http.ResponseWriter, r *http.Request, url *url.URL) {

	// Redirect with 302
	http.Redirect(w, r, url.EscapedPath(), http.StatusFound)
}

var r *mux.Router

func GetRouter() *mux.Router {

	if r == nil {

		r = mux.NewRouter()

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

		unlogged.HandleFunc("/login", LoginGet).Methods("GET").Name("login-get")
		unlogged.HandleFunc("/login", LoginPost).Methods("POST").Name("login-post")

		unlogged.HandleFunc("/register", RegisterGet).Methods("GET").Name("register-get")
		unlogged.HandleFunc("/register", RegisterPost).Methods("POST").Name("register-post")

		admin.HandleFunc("/addgenero", AddGeneroGet).Methods("GET").Name("addgenero-get")
		admin.HandleFunc("/addgenero", AddGeneroPost).Methods("POST").Name("addgenero-post")

		logged.HandleFunc("/logout", Logout).Name("logout")
	}

	return r
}
