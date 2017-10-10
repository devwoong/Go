package main

import (
	"Server/auth"
	"Server/example"
	"log"
	"net/http"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	// gomniauth.SetSecurityKey("PUT YOUR KEY")
	gomniauth.SetSecurityKey("woong")
	gomniauth.WithProviders(
		//facebook.New("key", "secret", "http://localhost:8080/auth/callback/facebook"),
		//github.New("key", "secret", "http://localhost:8080/auth/callback/github"),
		//google.New("key", "secret", "http://localhost:8080/auth/callback/google"),
		facebook.New("key", "secret", "http://localhost:8080/auth/callback/facebook"),
		github.New("key", "secret", "http://localhost:8080/auth/callback/github"),
		google.New("1085126168381-t281rb019442a0g486aerkdvg964q8fd.apps.googleusercontent.com",
			"6YQN9fKlDfTnvp7wyIcGfe3v", "http://localhost:8080/auth/callback/google"),
	)
	http.Handle("/login", &example.TemplateHandler{Filename: "login.html"})
	http.Handle("/default", auth.MustAuth(&example.TemplateHandler{Filename: "default.html"}))
	http.HandleFunc("/auth/", auth.LoginHandler)
	//	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("path/to/assets/"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
