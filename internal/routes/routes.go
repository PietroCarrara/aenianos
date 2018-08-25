package routes

import (
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request, url string) {

	// Redirect with 302
	http.Redirect(w, r, url, http.StatusFound)
}
