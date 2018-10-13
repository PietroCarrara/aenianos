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

	content := []byte(templates.Login(ctx))

	ctx.Close()

	w.Write(content)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

	user := r.PostFormValue("usuario")
	pass := r.PostFormValue("senha")

	u := aenianos.User{}
	data.Db.Where("username = ?", user).First(&u)

	if u.ID > 0 {

		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
			// Errou a senha
			// TODO: Avisar do erro
		} else {

			url := util.GetURL("user-get")
			defer util.Redirect(w, r, url)

			ctx := context.GetContext(w, r)
			defer ctx.Close()

			ctx.Session.Values["User.ID"] = u.ID

		}
	} else {
		// Errou o usuário
		// TODO: Avisar do erro
	}
}
