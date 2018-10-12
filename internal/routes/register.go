package routes

import (
	"log"
	"net/http"

	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/context"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/PietroCarrara/aenianos/internal/templates"
	"golang.org/x/crypto/bcrypt"
)

func RegisterGet(w http.ResponseWriter, r *http.Request) {

	ctx := context.GetContext(w, r)

	content := []byte(templates.Register(ctx))

	ctx.Close()

	w.Write(content)
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

	url, _ := GetRouter().Get("user-get").URL()
	defer Redirect(w, r, url)

	ctx := context.GetContext(w, r)
	defer ctx.Close()

	ctx.Session.Values["User.ID"] = u.ID
}
