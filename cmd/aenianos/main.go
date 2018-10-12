//
// main.go
// Copyright (C) 2018 pietro <pietro@the-arch>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"net/http"
	"os"

	"github.com/PietroCarrara/aenianos/internal/routes"
)

func main() {

	r := routes.GetRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, r)
}
