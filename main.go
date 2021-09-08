package main

import (
	"fmt"
	"net/http"

	"github.com/gera9/lenslocked.com/controllers"
	"github.com/gorilla/mux"
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 - <b>Page not found</b>")
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.FAQ).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(pageNotFound)

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
