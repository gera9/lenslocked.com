package main

import (
	"fmt"
	"net/http"

	"github.com/gera9/lenslocked.com/views"
	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	failOnError(homeView.Render(w, nil))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	failOnError(signupView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	failOnError(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	failOnError(faqView.Render(w, nil))
}

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
	homeView = views.NewView("bootswatch", "views/home.gohtml")
	contactView = views.NewView("bootswatch", "views/contact.gohtml")
	faqView = views.NewView("bootswatch", "views/faq.gohtml")
	signupView = views.NewView("bootswatch", "views/signup.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/signup", signUp)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(pageNotFound)

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
