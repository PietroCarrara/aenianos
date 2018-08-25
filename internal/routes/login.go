package routes

import (
	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/PietroCarrara/aenianos/internal/templates"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {

	ctx := context.GetContext(r)

	w.Write([]byte(templates.Login(ctx)))
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

			sess, err := data.Store.Get(r, data.MainSession)
			if err != nil {
				log.Fatal(err)
			}

			sess.Values["User.ID"] = u.ID
			sess.Save(r, w)

			Redirect(w, r, "/user")
		}
	} else {
		// Errou o usuário
		// TODO: Avisar do erro
	}
}
