package main

import (
	"fmt"
	"forum/internal/db"
	"forum/internal/handlers"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	db.Migrations()

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/auth", handlers.AuthPageHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
