package routes

import (
	"net/http"

	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/PietroCarrara/aenianos/internal/templates"
	"github.com/PietroCarrara/aenianos/internal/util"
	"golang.org/x/crypto/bcrypt"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {

	ctx := context.GetContext(w, r)

	redirect := r.FormValue("redirect")

	content := []byte(templates.Login(ctx, redirect))

	ctx.Close()

	w.Write(content)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

	user := r.PostFormValue("usuario")
	pass := r.PostFormValue("senha")
	redirect := r.PostFormValue("redirect")

	ctx := context.GetContext(w, r)

	u := aenianos.User{}
	data.Db.Where("username = ?", user).First(&u)

	if u.ID > 0 {

		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
			// Errou a senha
			// TODO: Avisar do erro
		} else {

			if redirect == "" {
				redirect = util.GetURL("user-get")
			}

			ctx.User = &u

			defer util.Redirect(w, r, redirect)
		}
	} else {
		ctx.AddFlash("Usuário ou senha incorretos!")

		url := util.GetURL("user-get")
		if redirect != "" {
			url += "?redirect=" + redirect
		}
		defer util.Redirect(w, r, url)
	}
	ctx.Close()
}
