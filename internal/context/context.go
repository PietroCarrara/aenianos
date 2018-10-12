package context

import (
	"log"
	"net/http"

	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/gorilla/sessions"
)

type Context struct {
	User    *aenianos.User
	Session *sessions.Session

	w http.ResponseWriter
	r *http.Request
}

func (c *Context) Close() {

	c.Session.Save(c.r, c.w)
}

func (c *Context) Closee(w http.ResponseWriter, r *http.Request) {

	c.Session.Save(r, w)
}

func GetContext(w http.ResponseWriter, r *http.Request) Context {

	res := Context{w: w, r: r}

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
