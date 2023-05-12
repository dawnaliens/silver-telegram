package main

import (
	"fmt"
	"github.com/dawnaliens/silver-telegram.git/controllers"
	"github.com/dawnaliens/silver-telegram.git/templates"
	"github.com/dawnaliens/silver-telegram.git/views"
	"github.com/go-chi/chi/v5"
	_ "html/template"
	"net/http"
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

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
