package routes

import (
	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/templates"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	ctx := context.GetContext(r)

	w.Write([]byte(templates.Hello(ctx)))
}
