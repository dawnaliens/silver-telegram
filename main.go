package main

import (
	"fmt"
	"github.com/dawnaliens/silver-telegram.git/controllers"
	"github.com/dawnaliens/silver-telegram.git/views"
	"github.com/go-chi/chi/v5"
	"net/http"
	"path/filepath"
)

//func executeTemplate(w http.ResponseWriter, filepath string) {
//	t, err := views.Parse(filepath)
//	if err != nil {
//		log.Printf("parsing template: %v", err)
//		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
//		return
//	}
//	t.Execute(w, nil)
//}

//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	tplPath := filepath.Join("templates", "home.gohtml")
//	executeTemplate(w, tplPath)
//
//}

//func contactHandler(w http.ResponseWriter, r *http.Request) {
//	tplPath := filepath.Join("templates", "contact.gohtml")
//	executeTemplate(w, tplPath)
//}
//
//func questionHandler(w http.ResponseWriter, r *http.Request) {
//	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
//}

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
	r := chi.NewRouter()
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))
	//var router http.HandlerFunc
	//router = pathHandler
	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on: 3000....")
	http.ListenAndServe(":3000", r)
}
