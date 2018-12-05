package main

import (
	"fmt"
	"gobots/controller"
	"gobots/route"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	Serverconfig := controller.ServerConfig()

	handler := http.NewServeMux()
	handler.HandleFunc("/register/", register)
	s := http.Server{
		Addr:         Serverconfig.Port,
		Handler:      handler,
		ReadTimeout:  time.Duration(Serverconfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(Serverconfig.WriteTimeout) * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

func register(responseW http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		t, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(responseW, err.Error(), 400)
			return
		}
		if err := t.Execute(responseW, nil); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		route.Register(responseW, request)
		http.Redirect(responseW, request, "/home/", 302)
	}

}
