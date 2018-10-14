package routes

import (
	"net/http"

	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/util"
)

func AccessLevel(level int) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.GetContext(w, r)

			if ctx.User == nil {
				// User is not logged. Go to login
				url := util.GetURL("login-get")
				redirect := r.URL.EscapedPath()
				util.Redirect(w, r, url+"?redirect="+redirect)
				return
			}

			if ctx.User.AccessLevel >= level {
				next.ServeHTTP(w, r)
			} else {
				// User is not authorized.
				// TODO: add errors to the session
				url := util.GetURL("home-get")
				util.Redirect(w, r, url)
			}
		})
	}

}

func Unlogged(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.GetContext(w, r)

		if ctx.User == nil {
			next.ServeHTTP(w, r)
		} else {
			// user is logged in
			url := util.GetURL("home-get")
			util.Redirect(w, r, url)
		}
	})
}
