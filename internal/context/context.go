package context

import (
	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

type Context struct {
	User    *aenianos.User
	Session *sessions.Session
}

func GetContext(r *http.Request) Context {

	res := Context{}

	// Get session
	sess, err := data.Store.Get(r, data.MainSession)
	if err != nil {
		log.Fatal(err)
	}
	res.Session = sess

	// Get logged user
	id, ok := sess.Values["User.ID"]
	if ok {
		u := aenianos.User{}
		data.Db.First(&u, id.(uint))
		res.User = &u
	}

	return res
}
