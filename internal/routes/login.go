package routes

import (
	"github.com/PietroCarrara/aenianos/internal/templates"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(templates.Login()))
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

}
