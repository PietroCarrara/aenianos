package middleware

import (
	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/data"
	"log"
	"net/http"
)

func AccessLevel(level int, route func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		sess, err := data.Store.Get(r, data.MainSession)
		if err != nil {
			log.Fatal(err)
		}
		id, ok := sess.Values["User.ID"]
		if !ok {
			// user is not logged in
			// TODO: redirect to login
			return
		}

		user := aenianos.User{}
		data.Db.First(&user, id.(uint))
		if user.ID <= 0 {
			log.Fatal("Invalid User.ID in session!")
		}

		if user.AccessLevel >= level {
			route(w, r)
		} else {
			// TODO: redirect somewhere
		}
	}
}

func Unlogged(route func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		sess, err := data.Store.Get(r, data.MainSession)
		if err != nil {
			log.Fatal(err)
		}
		_, ok := sess.Values["User.ID"]
		if !ok {
			route(w, r)
		} else {
			// user is logged in
			// TODO: redirect to home
		}
	}
}
