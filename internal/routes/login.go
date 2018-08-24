package routes

import (
	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/PietroCarrara/aenianos/internal/templates"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(templates.Login()))
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

	user := r.PostFormValue("usuario")
	pass := r.PostFormValue("senha")

	u := aenianos.User{}
	data.Db.Where("username = ?", user).First(&u)

	if u.ID > 0 {

		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
			// Errou a senha
		} else {

			sess, err := data.Store.Get(r, data.MainSession)
			if err != nil {
				log.Fatal(err)
			}

			sess.Values["User.ID"] = u.ID
			sess.Save(r, w)

			// TODO: redirect
		}
	} else {
		// Errou o usuário
	}
}
