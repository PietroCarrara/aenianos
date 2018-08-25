package routes

import (
	"github.com/PietroCarrara/aenianos/internal/data"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	sess, err := data.Store.Get(r, data.MainSession)

	if err != nil {
		log.Fatal(err)
	}

	// Reset all values
	sess.Values = map[interface{}]interface{}{}

	sess.Save(r, w)

	Redirect(w, r, "/user")
}
