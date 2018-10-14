package context

import (
	"log"
	"net/http"

	"github.com/PietroCarrara/aenianos"
	"github.com/PietroCarrara/aenianos/internal/data"
	"github.com/gorilla/sessions"
)

const mainSession = "session"

var store = sessions.NewCookieStore([]byte("secret"))

type Context struct {
	User *aenianos.User

	session *sessions.Session
	w       http.ResponseWriter
	r       *http.Request
}

func (c *Context) Clear() {
	c.session.Values = map[interface{}]interface{}{}
}

func (c *Context) Close() {

	if c.User != nil {
		c.session.Values["User.ID"] = c.User.ID
	} else {
		delete(c.session.Values, "User.ID")
	}

	c.session.Save(c.r, c.w)
}

func (c *Context) Flashes() []interface{} {
	return c.session.Flashes()
}

func (c *Context) AddFlash(val interface{}) {
	c.session.AddFlash(val)
}

func GetContext(w http.ResponseWriter, r *http.Request) Context {

	res := Context{w: w, r: r}

	// Get session
	sess, err := store.Get(r, mainSession)
	if err != nil {
		log.Fatal(err)
	}
	res.session = sess

	// Get logged user
	id, ok := sess.Values["User.ID"]
	if ok {
		u := aenianos.User{}
		data.Db.First(&u, id.(uint))
		res.User = &u
	}

	return res
}
