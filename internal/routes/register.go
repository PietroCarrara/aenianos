package routes

import (
	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/PietroCarrara/aenianos/internal/templates"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegisterGet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(templates.Register()))
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {

	user := r.PostFormValue("usuario")
	pass := r.PostFormValue("senha")

	// Verificar se já há alguém com o nome solicitado
	usuario := aenianos.User{}
	data.Db.Where("username = ?", user).First(&usuario)
	if usuario.ID > 0 {
		log.Println("Usuário já tomado!")
		return
	}

	// Hashear a senha
	cryptedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	u := aenianos.User{Username: user, Password: string(cryptedPass), AccessLevel: aenianos.ACCESS_NORMAL}

	data.Db.Save(&u)

	sess, err := data.Store.Get(r, data.MainSession)
	if err != nil {
		log.Fatal(err)
	}

	sess.Values["User.ID"] = u.ID

	sess.Save(r, w)
}
