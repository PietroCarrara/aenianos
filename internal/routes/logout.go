package routes

import (
	"github.com/PietroCarrara/aenianos/internal/context"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	defer Redirect(w, r, "/user")

	ctx := context.GetContext(w, r)
	defer ctx.Close()

	// Reset all values
	ctx.Session.Values = map[interface{}]interface{}{}
}
