package routes

import (
	"net/http"

	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {

	ctx := context.GetContext(w, r)

	content := []byte(templates.Hello(ctx))

	ctx.Close()

	w.Write(content)
}
