package main

import (
	"fmt"
	"forum/internal/db"
	"forum/internal/handlers"
	"log"
	"net/http"
)

func ErrorpageRender(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusNotFound:
		http.ServeFile(w, r, "templates/errors/404.html")
	case http.StatusInternalServerError:
		http.ServeFile(w, r, "templates/errors/500.html")
	case http.StatusBadRequest:
		http.ServeFile(w, r, "templates/errors/400.html")
	default:
		http.Error(w, http.StatusText(status), status)
	}
}

func main() {
	db.InitDB()
	defer db.CloseDB()

	db.Migrations()

	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/auth", handlers.AuthPageHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/home", handlers.HomeHandler)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.URL.Path != "/" {
	// 		ErrorpageRender(w, r, http.StatusNotFound)
	// 		return
	// 	}
	// 	handlers.HomeHandler(w, r)
	// })

	fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
