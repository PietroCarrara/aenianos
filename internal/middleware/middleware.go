package middleware

import (
	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/routes"
	"net/http"
)

func AccessLevel(level int, route func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.GetContext(w, r)
		defer ctx.Close()

		if ctx.User == nil {
			// User is not logged. Go to login
			routes.Redirect(w, r, "/login")
			return
		}

		if ctx.User.AccessLevel >= level {
			route(w, r)
		} else {
			// User is not authorized.
			// TODO: add errors to the session
			routes.Redirect(w, r, "/home")
		}
	}
}

func Unlogged(route func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.GetContext(w, r)
		defer ctx.Close()

		if ctx.User == nil {
			route(w, r)
		} else {
			// user is logged in
			routes.Redirect(w, r, "/home")
		}
	}
}
