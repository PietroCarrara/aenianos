package routes

import (
	"fmt"
	"net/http"

	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/PietroCarrara/aenianos/internal/templates"
	"github.com/PietroCarrara/aenianos/internal/util"
)

func AddGeneroGet(w http.ResponseWriter, r *http.Request) {

	ctx := context.GetContext(w, r)

	content := []byte(templates.AddGenero(ctx))

	ctx.Close()

	w.Write(content)
}

func AddGeneroPost(w http.ResponseWriter, r *http.Request) {

	fmt.Println("AAA")

	url := util.GetURL("addgenero-get")
	defer util.Redirect(w, r, url)

	ctx := context.GetContext(w, r)
	defer ctx.Close()

	nome := r.PostFormValue("nome")

	var genDb aenianos.Genero
	data.Db.Where("LOWER(nome) = LOWER(?)", nome).First(&genDb)
	if genDb.ID != 0 {
		ctx.Session.AddFlash("Nome já registrado!")
	} else {
		gen := aenianos.Genero{Nome: nome}

		data.Db.Save(&gen)
	}
}
