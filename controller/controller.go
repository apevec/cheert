// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/apevec/cheert/controller/about"
	"github.com/apevec/cheert/controller/home"
	"github.com/apevec/cheert/controller/static"
	"github.com/apevec/cheert/controller/up"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	//	debug.Load()
	//	register.Load()
	//	login.Load()
	home.Load()
	static.Load()
	//	status.Load()
	//	notepad.Load()
	up.Load()
}
