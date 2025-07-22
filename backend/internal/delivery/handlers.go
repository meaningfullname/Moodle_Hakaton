package delivery

import "net/http"

func handlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", SignIn)

	return mux
}
