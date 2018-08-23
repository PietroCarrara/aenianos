package routes

import (
	"net/http"
)

var routes map[string]func(http.ResponseWriter, *http.Request)

func init() {

	routes = map[string]func(http.ResponseWriter, *http.Request){}

	routes["/home"] = home

}

func Routes() map[string]func(http.ResponseWriter, *http.Request) {
	return routes
}
