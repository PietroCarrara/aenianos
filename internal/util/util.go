package util

import (
	"net/http"

	"github.com/PietroCarrara/aenianos/internal/data"
)

func GetURL(route string, params ...string) string {

	url, _ := data.GetRouter().Get(route).URL(params...)

	return url.EscapedPath()
}

func Redirect(w http.ResponseWriter, r *http.Request, url string) {

	// Redirect with 302
	http.Redirect(w, r, url, http.StatusFound)
}
