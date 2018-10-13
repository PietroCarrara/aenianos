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

	"github.com/PietroCarrara/aenianos/internal/data"
	_ "github.com/PietroCarrara/aenianos/internal/routes" // Initialize the routes
)

func main() {

	r := data.GetRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, r)
}
