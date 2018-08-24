package routes

import (
	"github.com/PietroCarrara/aenianos/internal/templates"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(templates.Hello()))
}
