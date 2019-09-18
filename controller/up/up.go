// Package up
package up

import (
	"net/http"
	"strings"

	"github.com/apevec/cheert/lib/flight"
	"github.com/apevec/cheert/model/up"

	"github.com/blue-jay-fork/core/router"
)

// Load the routes.
func Load() {
	router.Post("/up", Index2)
	router.Post("/up/submit", Submit)
	router.Get("/up/submit", End)
}

// Index displays the home page.
func Index2(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("up/index")

	dfgs := []string{
		"ds-core",
		"cs-core",
		"sssd",
		"freeipa",
		"freeipa-ds",
		"freeipa-cs-in-ipa",
		"crypto",
		"audit-subsystem",
		"selinux-solutions",
		"security-compliance",
		"security-technologies",
		"common-logging-and-session-recording",
		"openstack-identity-and-security-engineering",
		"special-projects",
		"special-projects-and-integration",
		"samba",
	}

	var dfgComma string
	for _, dfg := range dfgs {
		if strings.Contains(r.FormValue("group"), dfg) {
			dfgComma += dfg + ","
		}
	}

	v.Vars["group"] = strings.TrimSuffix(dfgComma, ",")
	v.Vars["id"] = r.FormValue("id")
	v.Vars["uid"] = r.FormValue("uid")

	v.Render(w, r)
}

func Submit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	r.ParseForm()
	for i, value := range r.Form["uid"] {
		plus := "0"
		if _, ok := r.Form["plus["+value+"]"]; ok {
			plus = r.Form["plus["+value+"]"][0] //do something here
		}

		_, err := up.Create(c.DB, r.Form["from"][0], value, plus, r.Form["praise"][i])
		if err != nil {
			c.FlashErrorGeneric(err)
			return
		}

	}

}

func End(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("up/submit")

	v.Render(w, r)
}
