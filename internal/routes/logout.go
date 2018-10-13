package routes

import (
	"net/http"

	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/util"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	url := util.GetURL("user-get")
	defer util.Redirect(w, r, url)

	ctx := context.GetContext(w, r)
	defer ctx.Close()

	// Reset all values
	ctx.Session.Values = map[interface{}]interface{}{}
}
