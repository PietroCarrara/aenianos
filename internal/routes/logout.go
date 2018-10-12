package routes

import (
	"net/http"

	"github.com/PietroCarrara/aenianos/internal/context"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	url, _ := GetRouter().Get("user-get").URL()
	defer Redirect(w, r, url)

	ctx := context.GetContext(w, r)
	defer ctx.Close()

	// Reset all values
	ctx.Session.Values = map[interface{}]interface{}{}
}
