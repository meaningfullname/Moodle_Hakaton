package main

import (
	"log"
	"net/http"

	_ "moodle/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// mux.HandleFunc("/api/v1/hello", helloHandler)

	log.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", mux)
}

// func ErrorRespone(status int, err error) {

// 	resp := struct{
// 		Status	int		`json:"status"`
// 		Error	error	`json:""`
// 	}{

// 	}
// 	_, err := json.Marshal()
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	w.Write([]byte("Hello World!"))
// }
