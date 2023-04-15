package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a"+
		"href=\"mailto:richardliux23@gmail.com\">richardliux23@gmail.com</a>")
}

func questionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ page</h1>")
}

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	default:
//		http.Error(w, "Page not found", http.StatusNotFound)
//		//w.WriteHeader(http.StatusNotFound)
//		//fmt.Fprint(w, "Page not found")
//	}
//}

func main() {
	//var router http.HandlerFunc
	//router = pathHandler
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", questionHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on: 3000....")
	http.ListenAndServe(":3000", r)
}
