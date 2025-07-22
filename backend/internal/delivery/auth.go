package delivery

import "net/http"

func SignIn(w http.ResponseWriter, r *http.Request) {
	log := r.FormValue("login")
	pwd := r.FormValue("password")

}
