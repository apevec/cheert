// Package home displays the Home page.
package home

import (
	"net/http"

	"github.com/apevec/cheert/lib/flight"

	"github.com/blue-jay-fork/core/router"
)

// Load the routes.
func Load() {
	router.Get("/", Index)
}

// Index displays the home page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("home/index")
	v.Render(w, r)
}
