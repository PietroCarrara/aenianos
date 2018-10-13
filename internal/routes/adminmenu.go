package routes

import (
	"net/http"

	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/templates"
)

func AdminMenu(w http.ResponseWriter, r *http.Request) {

	ctx := context.GetContext(w, r)

	content := templates.AdminMenu(ctx)

	ctx.Close()

	w.Write([]byte(content))
}
